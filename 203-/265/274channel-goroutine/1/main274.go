/*
问题:
完成goroutine和channel协同工作的案例,具体要求如下:
1.开启一个writeData协程,向管道intChan中写入50个整数
2.开启一个readData协程,从管道intChan中读取writeData写入的数据
3.注意:writeData和readData是同一个管道
4.主线程需要等待writeData和readData协程都完成工作才能退出
*/
package main

import (
	"fmt"
	"time"
)

func write(p chan int) {

	for i := 0; i < 50; i++ {
		p <- i
	}
}
func read(p chan int) {
	close(p)
	for v := range p {
		fmt.Println("v=", v)
	}
}
func main() {
	intChan := make(chan int, 50)
	go write(intChan)
	go read(intChan)
	time.Sleep(time.Second * 100)
}
