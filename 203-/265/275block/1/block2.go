package main

import (
	"fmt"
	"time"
)

// write Data
func writeData(p chan int) {
	for i := 1; i <= 1000; i++ {
		// 主函数中传入的管道intChan的容量是5,达到容量后会阻塞,除非读操作取出管道中的数据
		//放入数据
		p <- i
		fmt.Printf("writeData 写入数据=%v\n", i)
		println()
		time.Sleep(time.Second)
	}
	close(p) //关闭
}

// read data
func readData(p chan int) {
	for {
		v, ok := <-p //读取管道数据,每读一次管道数量少1
		if !ok {
			break
		}
		fmt.Printf("管道的长度=%v\n", len(p))
		fmt.Printf("readData 读到数据=%v\n", v)
		println()
		time.Sleep(time.Second * 3)
	}

}

// 总结:如果编译器运行过程中,发现一个管道只有写没有读(管道的长度一直增加),则该管道会堵塞
// 如果读和写都有,但是读写频率不一致,不会堵塞
func main() {
	//创建两个管道
	intChan := make(chan int, 1000)
	go writeData(intChan) //启动协程
	go readData(intChan)  //启动协程
	// time.Sleep(time.Second * 6)
	// 非阻塞接收并打印数据
	for i := 0; i < 1000; i++ {
		select {
		case data, ok := <-intChan:
			if ok {
				fmt.Printf("Received 接收数据=%v\n", data)
				println()
			} else {
				fmt.Println("Channel is closed")
			}
		default:
			fmt.Println("No data available")
		}
		time.Sleep(time.Second * 2)
	}
}
