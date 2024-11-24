package main

import "fmt"

// 作业6. 定义一个小计算器结构体(calcuator),实现加减乘除四个功能；实现方式1：分四个方法完成；实现方法2：用一个方法完成
// 作业6思路，如果需要分四个方法解决，粉笔计算+-*/;如果一个方法解决，需要方法参数中接收两个数和一个运算符
type Cal2 struct {
	Num1 float64
	Num2 float64
}

func (c1 *Cal2) getSum() float64 { //用指针传速度更快
	return c1.Num1 + c1.Num2 //直接调用结构体中的字段
}
func (c1 *Cal2) getSub() float64 {
	return c1.Num1 - c1.Num2
}
func (c1 *Cal2) getMul() float64 {
	return c1.Num1 * c1.Num2
}
func (c1 *Cal2) getDiv() float64 {
	return c1.Num1 / c1.Num2
}

func main() {
	var p Cal2
	p.Num1 = 1.2
	p.Num2 = 2.4
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", p.getSum()))
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", p.getSub()))
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", p.getMul()))
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", p.getDiv()))
}
