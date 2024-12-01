/*
1.获取1--100万之间的的素数,由于数量较多,用多个协程,此处用12个协程
2.退出通道中的容量也设置成12,,退出机制是退出通道exitChan被写满退出,
3.然后触发捕获协程退出信号,退出后关闭primeChan
*/
package main

import (
	"fmt"
	"time"
)

// 向intChan放入1--1000000个数
func putNum(intChan chan int) {
	for i := 1; i <= 1000000; i++ {
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
	intChan := make(chan int, 1000000)
	primeChan := make(chan int, 1000000) //放入结果
	exitChan := make(chan bool, 12)      //12个协程 标识退出管道
	start := time.Now().Unix()
	go putNum(intChan) //开启一个协程,向intChan放入1--8000个数
	//开启12个协程,从intChan取出数据并判断是否为素数,如果是就放入到primeChan
	for i := 0; i < 12; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//主线程要进行处理
	go func() {
		/*
			启动后，进入for循环，等待从 exitChan 通道中读取 12 个 true 信号。
			读取到12个 true 信号后，关闭 primeChan 通道
		*/
		for i := 0; i < 12; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		//当我们从exitChan取出4个结果,证明4个协程全部执行完成,就可以放心的关闭primeChan
		close(primeChan)
	}() //匿名函数格式结尾有(),匿名函数不用外部定义,直接使用
	//遍历我们的primeNum,把结果取出

	for {
		res, ok := <-primeChan
		if !ok { //表示通道已经关闭,且数据已经读完没有可以读的数据,primeChan中的数据已经循环的读完了
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	// time.Sleep(1000 * time.Second)
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
