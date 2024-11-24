package main

import "fmt"

type MethodUtils2 struct {
}
type Cal struct {
	n1       float64
	n2       float64
	operator string
}

// 作业4. 编写方法判断一个数是奇数还是偶数。
// 作业5. 根据行、列、字符，打印对应行数和列数的字符，比如行3列5字符#，打印对应的效果
// 作业6. 定义一个小计算器结构体(calcuator),实现加减乘除四个功能；实现方式1：分四个方法完成；实现方法2：用一个方法完成
// 作业6思路，如果需要分四个方法解决，分别计算+-*/;如果一个方法解决，需要方法参数中接收两个数和一个运算符
func (mu *MethodUtils2) way(n int) {
	if n%2 == 0 {
		fmt.Println("您输入的是偶数")
	} else {
		fmt.Println("您输入的是奇数")
	}
}
func (mu MethodUtils2) Print3(m, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}
func (mu Cal) Print4(a, b float64, c string) float64 {
	var res float64
	switch c {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}
func main() {
	var p MethodUtils2
	var m int
	fmt.Println("输入一个整数判断它是奇数还是偶数")
	fmt.Scanln(&m)
	p.way(m)
	var x, y int
	var str string
	fmt.Println("分别输入行数、列数和字符串")
	fmt.Scanln(&x, &y, &str)
	p.Print3(x, y, str)
	var w Cal
	d, e, f := w.n1, w.n2, w.operator
	fmt.Println("分别输入两个数和一个加减乘除运算符")
	fmt.Scanln(&d, &e, &f)
	z := w.Print4(d, e, f)
	fmt.Println("运算结果是", z)

}
