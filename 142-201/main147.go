package main

import "fmt"

func test01(arr [3]int) { //[3]int [4]int都是数组类型，但不是同一个类型因为长度不一样，同时如果[]中没有数字，则是切片类型
	arr[0] = 88 //将第一个值改成88，默认是0
}

func test02(arr *[3]int) { //变成了指向数组的指针
	fmt.Printf("外部函数arr指针的地址=%p\n", &arr)
	(*arr)[0] = 88 //通过数组指针找到数组元素
}
func main() {
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	arr01[2] = 50
	// 1. 数组是多个相同类型数据的组合，一个数组一旦声明/定义，其长度是固定的，不能动态变化
	// 2. var arr []int，这时候arr就是一个slice切片
	// 3. 数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
	// arr01[2] = 50.0 //这里会报错，数据类型要一样
	// arr01[3] = 100  //长度是固定的，不能动态变化，否则报越界错误
	fmt.Println(arr01)
	// 4. 数组创建后，如果没有赋值，有默认值(零值)，
	// 数值类型数组：默认值0；字符串数组：默认值""; bool数组：默认值为false
	var arr1 [3]float32
	var arr2 [3]string
	var arr3 [3]bool
	fmt.Printf("arr01=%v,arr02=%v,arr03=%v\n", arr1, arr2, arr3)
	// 5. 使用数组的步骤1.声明数组并开辟空间 2. 给数组元素赋值(默认零值) 3.使用数组
	// 6. 数组的下标是从0
	// var arr04 [3]string //0--2
	// var index int = 3
	// arr04[index] = "tom" //因为下标是0--2，因此arr04[3]就越界
	// fmt.Println(arr04[index])
	// 7. 数组下标必须在指定范围内使用，否则报panic:数组越界，比如var arr[5]int 则有效下标为0--4
	// 8. go的数组默认属于值类型(其他编程语言可能是引用类型)，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println(arr) //即使test01栈中arr[0]中变成88，但是不影响main主函数中arr[0]值
	// fmt.Println(test02(arr))
	// 9. 如果想在其他函数中，去修改原来的数组，可以使用引用传递(指针传递)
	arr123 := [3]int{11, 22, 33}
	fmt.Printf("主函数中arr123的地址=%p\n", &arr123)
	p := &arr123
	test02(p)           //或者直接test02(&arr123)
	fmt.Println(arr123) //通过指针传递修改了arr123中的下标0的值为88
	// 10. 长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度，不一致则会传递错误
}
