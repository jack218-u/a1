package main

import "fmt"

func main() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] //20,30,40
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v", i, slice[i])
	}
	fmt.Println()
	//使用for-range方式遍历
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
}
