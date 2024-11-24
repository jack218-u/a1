package main

import "fmt"

func fb(n int) []uint64 {
	//声明一个切片，切片大小n
	fbnSlice := make([]uint64, n)
	//第一个和第二个数的斐波那契数为1，即下标为0和下标为1的元素为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//进行for循环来存放斐波拉契的数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
func main() {
	//斐波拉契数列
	//1. 可以接收一个n int
	//2. 能够将斐波拉契数列放到切片中
	//3. 提示，斐波拉契的数列形式
	// arr[0] = 1; arr[1] = 1; arr[2] = 2; arr[3] = 3; arr[4] = 5; arr[5] = 8
	//思路：
	//1. 声明一个函数fbn(n int) ([]uint64)
	//2. 编程fbn(n int) 进行for循环来存放斐波拉契的数列
	for {
		var n int
		fmt.Print("输入斐波拉契数  ") //需要大于等于2
		fmt.Scanln(&n)
		fbnSlice := fb(n)
		fmt.Println("返回的斐波拉契数列为", fbnSlice)
	}
}
