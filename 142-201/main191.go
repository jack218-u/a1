// 方法基本介绍
// Golang中的方法是作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型都可以有方法，而不仅仅是struct
package main

import "fmt"

type Person3 struct {
	Name string
}
type Monster4 struct {
}

func (p Person3) test() { //给Person3结构体类型绑定一个方法
	fmt.Println("test() name=", p.Name) //然后输出了一个Person3结构体的一个字段值
	p.Name = "jack"
	fmt.Println("test() name修改后=", p.Name)
}

//以上语法解析：
//1. func (p Person3) test() {}表示Person3结构体有一方法，方法名为test
//2. (p Person3)体现test方法是和Person3类型绑定的
//3. a.test()，表示通过a这个Person3结构体变量调用外部test方法，a在主函数中声明了，因此可以调用，否则会报错
func main() {
	var a Person3
	a.Name = "tom"
	a.test()                              //调用方法
	fmt.Println("main() a.Name=", a.Name) //输出的是tom，因为结构体是值传递，外部函数字段改变不会影响主函数,除非你传地址(指针)
	//总结：
	//1. test方法和Person3绑定
	//2. test方法只能通过Person3结构体类型的变量调用，不能直接调用，也不能使用其他类型变量调用
	//以下调用方式都是错误的
	// test() //直接调用会报错，必须要绑定一个结构体类型的变量
	// var dog Monster4
	// dog.test() //test()用其他类型结构体变量比如dog调用，会报错，因为test()方法与Person3类型绑定，而dog不是Person3类型，因此不能调用
	//3. func (p Person3) test() {}	//其中p表示哪个Person3变量调用，这个p就是它的副本，和函数传参非常相似
	//4. p这个名字是用户指定，不是固定的，也可以改成其他
	// 快速入门
	//1. 给Person3结构体添加speak方法，输出xx是个好人
	a.speak()
	//2. 给Person3结构体添加jisuan方法，可以计算从1+...+1000的结果，说明方法内可以和函数一样进行各种运算
	a.jisuan()
	//3. 给Person3结构体添加jisuan2方法，该方法可以接收一个数n，计算从1+..+n的结果
	a.jisuan2(100) //n指定100，也可以修改任意值，则输出任意想要的结果
	//4. 给Person3结构体添加getSum方法，可以计算两个数的和，并返回结果；前面3个例子都没有要求有返回值
	b := a.getSum(100, 500) //getSSum()方法有返回值，但是没有在main主函数体中输出，所以定义变量b接收
	fmt.Println(b)
	//5. 方法的调用
	// a.speak()
	// a.jisuan()
	// a.jisuan2(100)
	// b := a.getSum(100, 500)
	// fmt.Println(b)
}
func (p Person3) speak() {
	fmt.Println(p.Name, "是个好人")
}
func (p Person3) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是=", res)
}
func (p Person3) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println("接收一个数n,计算1+...+n的结算结果", res)
}
func (p Person3) getSum(n1 int, n2 int) int {
	return n1 + n2 //返回去了，但是没有在自身的主函数体中输出，需要主函数体中定义变量接收
}
