package main

import "fmt"

var name = "tom" //全局变量

func test1() {
	name := "tom"
	fmt.Println(name)
}
func test2() {
	name = "jack"
	fmt.Println(name)
}
func main() {
	fmt.Println(name) //tom
	test1()           //tom
	test2()           //编译器采取就近原则，输出jack
	test1()           //tom
}
