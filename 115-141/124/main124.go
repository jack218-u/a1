package main

import "fmt"

//累计计数器
func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += string(36) //string(36)的字符为$
		fmt.Printf("str的值=%s ", str)
		return n
	}
}
func main() {
	a := AddUpper()
	fmt.Printf("a(1)的值=%v\n", a(1)) //11
	fmt.Printf("a(2)的值=%v\n", a(2)) //13
	fmt.Printf("a(3)的值=%v\n", a(3)) //16
}
