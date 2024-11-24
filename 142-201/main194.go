//方法的声明(定义)
/*
func (recevier type) methodName (参数列表) (返回值列表) {
方法体
return 返回值
}
*/

//1. recevier type:表示这个方法和type这个类型进行绑定,或者说该方法作用于type类型,receiver是这个方法的一个变量
//2. recevier type:type可以是结构体,也可以是其他的自定义类型
//3. receiver:就是type类型的一个变量(实例),比如:Person结构体的一个变量(实例)
//4. 参数列表:表示方法输入,类似于函数的形参列表,不需要可以为空
//5. 返回值列表:表示返回的值,可以多个,如果是单个可以不加括号,不写表示没有返回值
//6. 方法主体:表示为了实现某一功能的代码块
//7. return语句不是必须的,但是如果有返回值列表,必须要有return

// 方法的注意事项和细节讨论
// 1. 结构体类型是值类型,在方法调用中,遵守值类型的传递机制,是值拷贝传递
// 2. 如果需要修改结构体变量的值,可以通过结构体指针的方式来处理,即和结构体指针进行绑定
// 3. Golang中的方法作用在指定的数据类型上(即和指定的数据类型绑定),因此自定义类型,都可以有方法,而不仅仅是struct,比如int,float32等都可以有方法
// 4. 方法的访问范围控制规则,和函数一样.方法名首字母小写,只能在本包访问,方法名首字母大写,可以在本包和其他包访问
// 5. 如果一个类型实现了String()这个方法,那么fmt.Println默认会调用这个变量的String()进行输出
package main

import "fmt"

type Circle1 struct {
	radius float64 //radius为圆的半径
}

//为了提高效率,通常方法和结构体的指针类型绑定
func (c *Circle1) area2() float64 {
	//因为c是指针,因此标准的访问其字段方式是(*c).radius
	// return 3.14 * (*c).radius * (*c).radius
	fmt.Printf("方法中c是*Circle1指向的地址=%p\n", c) //c前面不需要加&,因为它本身就是指针,如果加&表示它本身的地址,不是指针地址
	c.radius = 10                            //外部变量的改变会影响主函数的变化
	return 3.14 * c.radius * c.radius

}
func main() {
	var c1 Circle1
	fmt.Printf("main主函数中的c1结构体变量地址=%p\n", &c1) //c1的指针地址和方法area2中的c地址一样
	c1.radius = 6.0
	// res := (&c1).area2()
	res := c1.area2()
	//编译器底层做了优化,(&c1).area2()等价于c1.area2()
	//因为编译器底层会自动给加上&
	fmt.Println("圆的面积=", res)
	fmt.Println("c1.radius=", c1.radius)
	var i integer = 10
	i.print()
	i.change() //引入指针改变了原来的i值
	fmt.Println("i=", i)
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(stu)
	fmt.Println(&stu) //如果你实现了 *Student类型的String方法,就会自动调用
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

//给Student实现方法String()方法
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}
