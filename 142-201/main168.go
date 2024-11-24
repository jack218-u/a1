package main

import "fmt"

func main() {
	//二位数组得演示案例
	// 0 0 0 0 0 0
	// 0 0 1 0 0 0
	// 0 2 0 3 0 0
	// 0 0 0 0 0 0
	//声明二位数组
	//1. 先声明再赋值，语法：var 数组名 [a1][a2]类型，a1表示二维数组中共有几个一维数组，a2表示每个一维数组中有几个元素,比如var arr [2][3]int
	//2. 数组赋值中元素赋值方法，比如，数组名[b1][b2],b1表示表示二维数组中排序为b1的一维数组（排序从0开始）
	//中的的排序为b2（排序为从0开始）的元素的值
	var arr [4][6]int
	//赋初值
	fmt.Println("输出没有赋值的数组", arr)
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println("输出赋值以后的数组", arr)
	//遍历二维数组，按照要求输出图形
	//方法1：
	fmt.Println("方法1输出")
	for _, v := range arr { //遍历数组
		for _, v2 := range v {
			fmt.Print(v2, " ") //每个输出后面需要加一个空格
		}
		fmt.Println() //每个内循环后需要换行
	}
	//方法2
	fmt.Println("方法2输出")
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ") //每个输出后面需要加一个空格
		}
		fmt.Println() //每个内循环后需要换行
	}
}
