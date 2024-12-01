package main

import "fmt"

func main() {
	mapChan := make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"
	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	mapChan <- m1
	mapChan <- m2
	/* 	使用 select 语句：
	如果你希望在多个 channel 之间进行选择，可以使用 select 语句。这在处理多个并发任务时非常有用
	*/
	// 使用 select 语句接收并打印数据
	for i := 0; i < 2; i++ {
		select {
		case data := <-mapChan:
			fmt.Println("Received:", data)
		default:
			fmt.Println("No data available")
		}
	}
	//这种方法会阻塞当前 goroutine，直到 channel 中有数据可读
	fmt.Printf("mapChan len=%d cap=%d\n", len(mapChan), cap(mapChan))
}
