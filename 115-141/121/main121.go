package main

import "fmt"

//用指针交换数据
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
func main() {
	a1 := 10
	b1 := 20
	swap(&a1, &b1) //main主函数中用指针变量调用外部函数变量后，可以改变主函数中的的变量数值
	fmt.Printf("a1的值=%v,b1的值=%v", a1, b1)
}
