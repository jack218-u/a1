package main

import "fmt"

type Cal4 struct {
}

func (arr Cal4) Print7(a [3][3]int) {
	// 创建一个新的 3x3 数组来存储转置后的结果
	var transposed [3][3]int

	// 用于存储每一行的字符串
	var rowStrings [3]string

	// 转置矩阵并构建每行的字符串
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[j][i] = a[i][j]
			rowStrings[j] += fmt.Sprintf("%v ", transposed[j][i])
		}
	}

	// 打印转置后的矩阵
	for _, row := range rowStrings {
		fmt.Println(row)
	}
}

func main() {
	var p Cal4
	var a [3][3]int
	a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
	fmt.Println()
	p.Print7(a)
}
