package main

import "fmt"

var (
	//Fun1就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//调用方式1：在定义匿名的时候就直接调用，这种方式匿名函数只能使用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20) //匿名函数使用的变量值分别为10，20
	fmt.Println("res1=", res1)
	//调用方式2：将匿名函数func (n1 int,n2 int) int 赋值给a
	//则a的数据类型就是函数类型,此时我们可以通过a完成调用
	//且可以多次调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2的值=", res2)
	res3 := a(90, 30) //通过变量a完成匿名函数调用，可以多次调用
	fmt.Println("res3的值=", res3)
	res4 := Fun1(10, 20)
	fmt.Println("res4的值=", res4)
}
