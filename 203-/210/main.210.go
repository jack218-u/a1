// 1. 结构体可以使用嵌套匿名结构体所有的字段和方法，即:首字母大写或者小写的字段、方法，都可以使用。
package main

import "fmt"

type A struct {
	Name string
	age  int //结构体可以使用嵌套匿名结构体所有的字段和方法，即:首字母大写或者小写的字段、方法，都可以使用。
}

type B struct {
	A
	Name string
}
type C struct {
	Name   string
	Scrore float64
}

type D struct {
	A
	C
}
type E struct {
	a A //有名结构体(组合关系)，名字为a指定的结构体是A，访问的时候必须带上a
}
type Goods struct {
	Name  string
	price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods //嵌套匿名结构体指针类型,返回的是地址，效率更高，访问的时候需要解引用字段
	*Brand
}

// 创建一个新的结构体来包含 *Goods 和 *Brand
type Combined struct {
	Goods Goods
	Brand Brand
}

func (b *B) SayOk() {
	fmt.Println("B SayOK", b.Name)
}
func (a *A) SayOk() {
	fmt.Println("A SayOK", a.Name)
}
func (a *A) hello() {
	fmt.Println("A SayOK", a.Name)
}
func (t *TV2) Print() string {
	if t == nil || t.Goods == nil || t.Brand == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Goods: %+v, Brand: %+v", *t.Goods, *t.Brand)
}
func main() {
	var b B
	//2. 匿名结构体字段访问可以简化
	// b.A.Name = "Tom" //可以简化成b.Name="Tom"
	// b.A.age = 19     //可以简化成b.age=19
	// b.A.SayOk()      //可以简写成b.SayOk()
	// b.A.hello()      //可以简写成b.hello()
	//以上代码小结：
	//a.当我们直接通过b访问字段或者方法时,其执行流程如下,比如b.Name
	//b.编译器会先看b对应的类型有没有Name,如果有,则直接调用B类型的Name字段
	//c.如果没有就去看B中嵌入的匿名结构体A有没有声明Name字段,如果有就调用,
	//如果没有就继续查找...(A中是否有匿名结构体,如此递归往下层查找),如果都找不到就报错
	//3. 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
	//   如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分，例如第42行
	b.Name = "jack" //  ok
	// b.A.Name = "scott"
	b.age = 19 //ok
	b.SayOk()  //B SayOK jack
	b.hello()  // b.hello()的时候，和b绑定的结构体B,没有对那个的hello方法，转向B中匿名结构体A，
	//A中有hello方法，对应的输出为a.Name，但是a.Name为空，所以输出为空,如果此时再加上b.A.Name = "scott"，
	//因为B中此时有自己的Name,所以b.A.Name = "scott",表示scott赋值给b中匿名结构体A中的Name在字段，结果为scott
	//4. 结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，
	//在访问时，就必须明确指定匿名结构体名字，否则编译错误。
	var d D
	// d.Name = "john" //此时会报错，因为d的类型D中有两个匿名结构体A和C中都有Name,不知道选择哪个
	d.A.Name = "john" //除非单独指定选择A.如左边例子，或者在D结构体中加入一个Name字段，就近原则选择
	fmt.Println(d.A.Name)
	//5. 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果组合关系，那么访问组合的结构体的字段活方法时，
	//必须带上结构体名字
	var e E
	e.a.Name = "henry" //访问方式为 变量+有名结构体字段+结构体名称
	fmt.Println(e.a.Name)
	//6. 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}} //必须按照顺序
	// 值嵌套
	tv2 := TV{
		Goods{
			price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)
	// 指针类型嵌套结构体赋值的时候，需要加&获取地址
	tv3 := TV2{&Goods{"电视机003", 7000.99}, &Brand{"创维", "深圳"}}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand) //解引用访问字段,同时分开取值,因为有两个匿名字段所以不能同时解引用输出
	// 指针类型嵌套结构体赋值的时候，需要加&获取地址
	tv4 := TV2{
		&Goods{
			price: 9000.99,
			Name:  "电视机004",
		},
		&Brand{
			Name:    "索尼",
			Address: "北京",
		},
	}
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand) // 解引用访问字段,同时分开取值,因为有两个匿名字段所以不能同时解引用输出
	fmt.Println("tv4", tv4.Print())
	tv4Combined := Combined{
		Goods: *tv4.Goods,
		Brand: *tv4.Brand,
	}
	fmt.Printf("tv4Combined: %+v\n", tv4Combined)
}
