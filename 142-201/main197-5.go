package main

import "fmt"

type Cal5 struct {
}

func (arr Cal5) Print8(a [4][4]int) {
	// 创建一个新的 4x4 数组来存储转置后的结果
	var transposed [4][4]int

	// 用于存储每一行的字符串
	var rowStrings [4]string

	// 转置矩阵并构建每行的字符串
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
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
	var p Cal5
	var a [4][4]int
	a = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
	fmt.Println()
	p.Print8(a)
}
