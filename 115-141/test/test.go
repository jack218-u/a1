package main

import "fmt"

var name = "tom" //全局变量

func test01() {
	fmt.Println(name)
}
func test02() {
	name = "jack"
	fmt.Println(name)
}
func main() {
	fmt.Println(name) //tom
	test01()          //tom
	test02()          //编译器采取就近原则，输出jack
	test01()          //tom
	fmt.Println(name)
}
