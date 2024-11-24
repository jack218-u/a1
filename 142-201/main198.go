// 方法和函数的区别
// 1. 调用方式不一样
// 函数调用方式：  函数名(实参列表)
// 方法的调用方式：变量.方法名(实参列表)

package main

import "fmt"

type Person4 struct {
	Name string
}

func test1(p Person4) {
	fmt.Println(p.Name)
}
func test2(p *Person4) {
	fmt.Println(p.Name)
}
func (p Person4) test3() {
	p.Name = "jack"
	fmt.Println("test3()=", p.Name)
}

func (p *Person4) test4() {
	p.Name = "mary"
	fmt.Println("test3()=", p.Name)
}
func main() {
	p := Person4{"tom"}
	// 2.对于普通函数，接收者为值类型，不能将指针类型的数据直接传递，反之亦然
	// test1(p)  //函数调用中接收者如果是值类型，则需要传递值类型
	// test2(&p) //函数调用中接收者如果是指针类型，则也需要传递指针类型

	// 3.1 对于方法(如struct方法)，接收者为值类型时，也可以用指针类型调用
	p.test3()
	fmt.Println("main中调用test3 p.name=", p.Name)
	(&p).test3() //只是形式上的地址拷贝，本质还是值拷贝，编译器做了优化处理
	//即此调用外部方法中的p.Name=jack，对于主函数中的p.name无影响，还是tom

	// 3.2 对于方法(如struct方法)，接收者为指针类型时，也可以用值类型调用
	(&p).test4()
	fmt.Println("main中调用test4 p.name=", p.Name)
	p.test4() //等价于 (&p).test4()
	fmt.Println("main中调用test4 p.name=", p.Name)
	//总结：因为方法是和类型(比如struct等)的变量绑定的，所以必须通过和方法绑定的类型的变量来调用，所以：
	//1.不管调用形式如何，最终决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定
	//2.如果是和值类型绑定，比如(p.Person)，则是值拷贝，如果和指针类型绑定，比如(p *Person)则是地址拷贝
}
