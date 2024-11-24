package main

import (
	"fmt"
	_ "math"
)

// 将计算功能放在一个函数中，需要使用时调用即可
func cal(n1 float64, n2 float64, operate byte) float64 { //返回值只有一个可以不加括号
	var res float64
	switch operate {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符好错误...")
	}
	return res
}
func main() {
	var n1, n2 float64 = 5.3, 2.3
	var operate byte = '-'
	result := cal(n1, n2, operate)
	fmt.Println("result=", result)
	n1 = 4.5
	n2 = 6.7
	operate = '*'
	result = cal(n1, n2, operate)
	fmt.Println("result=", result)
}
