package main

import "fmt"

func test(n int) {
	n = n + 1
	fmt.Println(n)
}
func main() {
	n := 20
	test(n)        //基本数据类型n（数组也是一样）在函数内修改，不影响原来的值，属于值传递
	fmt.Println(n) //所以n还是20
}