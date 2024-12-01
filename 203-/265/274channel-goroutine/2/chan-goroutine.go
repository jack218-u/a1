package main

import (
	"fmt"
	"time"
)

// write Data
func writeData(p chan int) {
	for i := 1; i <= 5; i++ {
		//放入数据
		p <- i
		fmt.Println("writeData", i)
		time.Sleep(time.Second)
	}
	close(p) //关闭
}

// read data
func readData(p chan int, q chan bool) {
	for {
		v, ok := <-p
		if !ok {
			break // 如果通道 p 已经关闭且没有数据可读取，则退出循环
		}
		fmt.Printf("readData 读到数据=%v\n", v)
		time.Sleep(time.Second)
	}
	//readData读取完数据后,即任务完成
	q <- true
	close(q)
}
func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			break // 如果 exitChan 已经关闭且没有数据可读取，则退出循环
		}
	}
}
