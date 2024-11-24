package main

import "fmt"

// 作业6. 定义一个小计算器结构体(calcuator),实现加减乘除四个功能；实现方式1：分四个方法完成；实现方法2：用一个方法完成
// 作业6思路，如果需要分四个方法解决，粉笔计算+-*/;如果一个方法解决，需要方法参数中接收两个数和一个运算符
type Cal3 struct {
	Num1 float64
	Num2 float64
}

func (c1 Cal3) getRes(c string) float64 {
	var res float64
	res = 0.0
	switch c {
	case "+":
		res = c1.Num1 + c1.Num2
	case "-":
		res = c1.Num1 - c1.Num2
	case "*":
		res = c1.Num1 * c1.Num2
	case "/":
		res = c1.Num1 / c1.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

func main() {
	var p Cal3
	p.Num1 = 1.2
	p.Num2 = 2.4
	res1 := p.getRes("+")
	res2 := p.getRes("-")
	res3 := p.getRes("*")
	res4 := p.getRes("/")
	fmt.Printf("res1=%.2f\n", res1)
	fmt.Printf("res2=%.2f\n", res2)
	fmt.Printf("res3=%.2f\n", res3)
	fmt.Printf("res4=%.2f\n", res4)
}
