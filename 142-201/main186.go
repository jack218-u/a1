package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

func main() {
	//创建结构体变量和访问结构体字段
	//方式1. 直接声明 var person struct {} //参考main183.go
	//方式2. var person Person = Person {}
	p2 := Person1{} //声明后，单独字段赋值
	p2.Name = "tom"
	p2.Age = 18
	fmt.Println(p2)
	p3 := Person1{"mary", 20} //直接给结构体字段赋值，最方便
	fmt.Println(p3)
	p8 := Person1{"张三", 25} //也可以直接给结构体字段赋值
	fmt.Println(p8)
	//方式3-&，var person *Person1 = new(Person1)，指向Person1的指针变量
	//因为p3是一个指针，因此标准的给字段赋值方式如下
	var p4 *Person1 = new(Person1)
	//(*p4).Name = "smith"也可以写成 p4.Name = "smith"
	//原因：go设计者为了使用方便，底层对p4.Name = "smith"进行处理，给p4加上取值运算 (*p4).Name = "smith"
	(*p4).Name = "smith"
	(*p4).Age = 30
	fmt.Println(*p4)
	var p5 *Person1 = new(Person1)
	p5.Name = "Amy"
	p5.Age = 30
	fmt.Println(*p5)
	//方式4-{}  var person *Person = &Person{},也可以在后面直接赋值var person *Person = &Person{字段值,字段值}
	var p6 *Person1 = &Person1{}
	// (*p6).Name = "scott" //标准写法
	p6.Name = "Brown" //简写，原因和方式3一样，在底层加上了*
	// (*p6).Age = 10
	p6.Age = 15
	fmt.Println(*p6)
	var p7 *Person1 = &Person1{"henry", 35}
	fmt.Println(*p7)
	//总结：
	//1. 第三种和第四种方式返回的是结构体指针
	//2. 结构体指针访问字段的标准方式应该是:(*结构体指针).字段名，比如(*p).Name = "tom"
	//3. 但go编程语言做了一个简化，也支持 结构体指针.字段名，比如 p.Name = "tom",更加符合使用习惯，go编译器底层对p.Name做了转换成(*p).Name

}
