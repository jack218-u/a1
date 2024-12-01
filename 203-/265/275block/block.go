package main

import (
	"fmt"
	"time"
)

// write Data
func writeData(p chan int) {
	for i := 1; i <= 100; i++ {
		// 主函数中传入的管道intChan的容量是5,达到容量后会阻塞,除非读操作取出管道中的数据
		//放入数据
		p <- i
		fmt.Println("writeData", i)
		// time.Sleep(time.Second)
	}
	close(p) //关闭
}

// read data
func readData(p chan int, q chan bool) {
	for {
		v, ok := <-p //读取管道数据,每读一次管道数量少1
		if !ok {
			break
		}
		fmt.Printf("管道的长度=%v\n", len(p))
		fmt.Printf("readData 读到数据=%v\n", v)
		time.Sleep(time.Second)
	}
	//readData读取完数据后,即任务完成
	q <- true
	close(q)
}

// 总结:如果编译器运行过程中,发现一个管道只有写没有读(管道的长度一直增加),则该管道会堵塞
// 如果读和写都有,但是读写频率不一致,不会堵塞
func main() {
	//创建两个管道
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 6)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
