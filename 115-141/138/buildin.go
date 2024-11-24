package main

import (
	"fmt"
)

func main() {
	//内置函数：1.len 2.new 3.make
	num1 := 100
	fmt.Printf("num1的类型是%T, num1的值=%v, num1的地址是%v\n", num1, num1, &num1)
	//new: 用来分配内存，主要分配值类型，比如int、float32、struct，返回的是指针
	num2 := new(int) // *int
	//num2的类型是%T => *int 整数指针
	//num2的值=%v  地址 （这个地址是系统分配）
	//num2的地址是%v  地址 （这个地址是系统分配）
	//num2这个指针指向的值
	*num2 = 100 //num2指针指向的值由于0 变成了 100
	fmt.Printf("num2的类型是%T, num2的值=%v, num2的地址是%v,num2这个指针指向的值=%v\n", num2, num2, &num2, *num2)
	//make主要分配引用类型的内存，比如chan、map、slice
}
