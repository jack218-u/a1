package main

import "fmt"

type Cal1 struct {
}

func (mul Cal1) Print5(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", j, i, j*i)
		}
		fmt.Println()
	}
}
func (arr Cal1) Print6(a [3][3]int) {
	// 创建一个新的 3x3 数组来存储转置后的结果
	var transposed [3][3]int
	// 转置矩阵
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//不能将 transposed[j][i] = a[i][j] 改为 a[i][j]，
			//因为这样会失去矩阵转置的功能，导致代码逻辑错误
			transposed[j][i] = a[i][j]
		}
	}
	// 打印转置后的矩阵
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", transposed[i][j])
		}
		fmt.Println()
	}
}
func main() {
	var p Cal1
	var n int
	fmt.Println("输入1-9,打印对应的乘法表")
	fmt.Scanln(&n)
	p.Print5(n)
	var a [3][3]int
	a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
	fmt.Println()
	p.Print6(a)
}
