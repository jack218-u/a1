package main

import "fmt"

func main() {
	// 1.定义一个数组
	var hens [7]float64
	//2.给数组的每个元素赋值,元素的下标是从0开始的 0--5
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0
	//3.遍历数组求组总体重
	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	//4.求出平均体重
	averageWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v,averageWeight2=%v\n", totalWeight2, averageWeight2)
}
