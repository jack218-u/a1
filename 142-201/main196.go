package main

import "fmt"

type MethodUtils struct {
}

//练习1,编写一个不需要任何参数的方法,在方法中打印一个2*3的矩形,注意方法对应的结构体可以没有任何字段
//练习2.编写一个方法,提供m和n两个参数,方法中打印一个m行n列的矩形
//练习3.编写一个方法算该矩形的面积(可以接收长m,宽n),将其作为方法返回值.在main方法中调用该方法,接收返回的面积值并打印
func (mu MethodUtils) Print() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mu MethodUtils) Print2(m, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mu MethodUtils) area(len, width float64) float64 {
	return len * width
}
func main() {
	var p MethodUtils
	p.Print()
	var m, n int
	var len, width float64
	fmt.Println("输入打印矩形的行数和列数:")
	fmt.Scanln(&m, &n)
	p.Print2(m, n)
	fmt.Println("输入矩形的长和宽:")
	fmt.Scanln(&len, &width)
	o := p.area(len, width)
	fmt.Println("矩形的面积是:", o)
}
