package main

import "fmt"

// func test02(n1 int) {
// 	n1 = n1 + 10
// 	fmt.Println("test02() n1= ", n1)
// }
func test03(n1 *int) {
	fmt.Printf("输出n1的地址%v\n", &n1) //n1虽然是指针但是也有自己本身的地址，输出方式为&n1
	*n1 = *n1 + 10
	// var ptr **int = &n1
	// fmt.Println("test03() n1-ptr= ", &(ptr)) //指针往下层传递，即指针数据自己的地址需要用中间变量传递
	fmt.Println("test03() n1= ", *n1)
}
func main() {
	// num := 20
	// test02(num)                      //基本数据类型n（数组也是一样）在函数内修改，不影响原来的值，属于值传递
	// fmt.Println("main() num= ", num) //所以n还是20
	num := 20
	fmt.Printf("输出num的地址%v\n", &num)
	// test02(num)
	test03(&num)
	fmt.Println("test03() n1= ", num)
}
