package main

import (
	"fmt"
	"sort"
)

func main() {
	//map排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)
	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key放入到切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key来输出map的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k) // 第一步，将map1中的key的值放入到切片keys中,map1的key是整数类型
	}
	//排序
	fmt.Println("排序前的切片", keys)
	sort.Ints(keys) //第二步，按照递增顺序对keys切片排序
	fmt.Println("排序后的切片", keys)
	fmt.Println("map1的排序排序后的结果为")
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k]) //第三步，遍历排序后的map的key值，输出排序后的map1
	}
}
