/*
需求:要求统计1-8000的数字中,哪些是素数,用goroutine和channel的知识后完成
分析思路:
1.传统方法,就是使用一个循环,循环的判断各个数是不是素数
2.使用并发/并行的方法,将统计素数的任务分配给多个(4个)goroutine去完成
*/

package main

import (
	"fmt"
	"time"
)

// 向intChan放入1--160000个数
func putNum(intChan chan int) {
	for i := 1; i <= 800000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

// 从intChan取出数据并判断是否为素数,如果是就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//管道类型数据的表示为chan + 数据类型,两个字符串表示
	//使用for循环
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { //intChan取不到...
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ { //此处i不能=num,否则此情况下num%i == 0,不符合预期
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数放入到primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据,退出")
	//这里还不能关闭primeChan
	//向exitChan写入true
	exitChan <- true //当前协程执行完毕,输出结果到exitChan管道,该管道容量4
}

func main() {
	intChan := make(chan int, 800000)
	primeChan := make(chan int, 800000) //放入结果
	//标识退出管道
	exitChan := make(chan bool, 16) //XX个协程
	//开启一个协程,向intChan放入1--8000个数
	start := time.Now().Unix()
	go putNum(intChan) //putNum一个协程
	//开启4个协程,从intChan取出数据并判断是否为素数,如果是就放入到primeChan
	for i := 0; i < 16; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//主线程要进行处理
	go func() {
		/*
			启动后，进入for循环，等待从 exitChan 通道中读取 4 个 true 信号。
			读取到4个 true 信号后，关闭 primeChan 通道
		*/
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		//当我们从exitChan取出4个结果,证明4个协程全部执行完成,就可以放心的关闭primeChan
		close(primeChan)
	}() //匿名函数格式结尾有(),匿名函数不用外部定义,直接使用
	//遍历我们的primeNum,把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			// 如果ok为true，表示通道仍然打开并且成功接收到数据；如果ok为false，表示通道已经关闭且没有更多的数据可以读取
			break
		}
		//将结果输出
		// fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main主线程退出")
}

/*
思维导图如下:
主程序入口 (main 函数)
├── 创建通道 intChan (缓冲大小 1000)
├── 创建通道 primeChan (缓冲大小 2000)
├── 创建通道 exitChan (缓冲大小 4)
├── 启动 putNum 协程
│   └── 循环 1 到 8000
│       └── 将数字发送到 intChan
│   └── 关闭 intChan
├── 启动 4 个 primeNum 协程
│   └── 持续从 intChan 读取数字
│   └── 判断数字是否为素数
│       └── 如果是素数，将数字发送到 primeChan
│   └── 当 intChan 关闭且无数据可读时，向 exitChan 发送 true
├── 启动匿名函数协程
│   └── 从 exitChan 读取 4 个 true
│   └── 关闭 primeChan
└── 主线程从 primeChan 读取所有素数并打印
    └── 打印 "main主线程退出"
*/
