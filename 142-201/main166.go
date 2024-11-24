package main

import "fmt"

//二分查找的函数
//二分查找思路：比如我们要查找的数是 findVal
//1. arr是一个有序数组，并且从小到大排序
//2. 先找到中间的下标middle = (leftIndex + rightIndex) / 2 然后让中间下标的值和findVal进行比较
//2.1 如果arr[middle] > findVal,就应该向leftIndex --- (middle - 1)
//2.2 如果arr[middle] < findVal,就应该向(middle + 1) --- rightIndex
//2.3 如果arr[middle] == findVal，就找到
//2.4 上面的2.1 2.2 2.3 的逻辑会递归执行
//3. 查找不到的情况---分析出退出递归的条件
// if leftIndex > rightIndex {
// 	找不到
// 	return
// }
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//说明我们要查找的数，应该在leftIndex ---middle-1，递归执行BinaryFind函数，
		//即middle := (leftIndex + rightIndex)/2 , 同时middle-1变成了最右边的下标
		BinaryFind(arr, leftIndex, middle-1, findVal) //因为前面middle下标的数不匹配，已经排查，所以右边界是middle-1，减少查询次数
	} else if (*arr)[middle] < findVal {
		//说明我们要找找的数，应该在 middle+1 ---rightIndex范围,递归执行BinaryFind函数，
		//即middle := (leftIndex + rightIndex)/2 , 同时middle+1变成了最左边的下标
		BinaryFind(arr, middle+1, rightIndex, findVal) //因为前面middle下标的数不匹配，已经排查，所以左边界是middle+1，减少查询次数
	} else {
		//找到了
		fmt.Printf("找到了,下标为%v \n", middle)
	}
}
func main() {
	arr := [6]int{1, 8, 10, 89, 1000}
	BinaryFind(&arr, 0, len(arr), 89)
}
