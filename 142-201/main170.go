package main

import "fmt"

func main() {
	//二维数组遍历
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//1.双层for循环
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}
	fmt.Println("打印二维数组arr3第一行的长度=", len(arr3[0])) //arr3的第一行的长度
	//2. for-range来遍历二维数组
	for i, v := range arr3 { //v在此处为一维数组
		for j, v2 := range v {
			fmt.Printf("arr3[%v][%v]=%v \t", i, j, v2)
		}
		fmt.Println()
	}
	fmt.Println("对二维数组arr3进行遍历")
	for i, value := range arr3 {
		fmt.Printf("i=%v value=%v ", i, value)
	}
}
