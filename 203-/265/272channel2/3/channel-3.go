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
	//使用非阻塞的方式尝试从 channel 中接收数据。如果 channel 中没有数据，程序不会阻塞。
	// 非阻塞接收并打印数据
	for i := 0; i < 20; i++ {
		select {
		case data, ok := <-mapChan:
			if ok {
				fmt.Println("Received:", data)
			} else {
				fmt.Println("Channel is closed")
			}
		default:
			fmt.Println("No data available")
		}
	}
	//这种方法会阻塞当前 goroutine，直到 channel 中有数据可读
	fmt.Printf("mapChan len=%d cap=%d\n", len(mapChan), cap(mapChan))
}
