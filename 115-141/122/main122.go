package main

import (
	"abc/function/utils"
	"fmt"
)

var age int = test()

// 为了看到全局变量是先被初始化，我们这里先写函数
func test() int {
	fmt.Println("test()") //1
	return 90
}

// init函数，通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...") //2
}
func main() {
	fmt.Println("main()...age", age) //3
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}
