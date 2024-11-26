/*
实现如下:
1.在主线程(可以理解成进程)中,开启一个goroutine,该协程每隔1s输出"hello world"
2.在主线程中也每隔1输出"hello golang",输出10s后结束
3.要求主线程和goroutine同时执行
4.画出主线程和协程执行流程图
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
1.主线程是一个物理线程,直接作用在cpu上,是重量级,非常消耗cpu;
2.协程是主线程开启的,是轻量级的线程,是逻辑态,对资源消耗相对小;
3.Golang的协程机制是重要的特点,可以轻松开启上万个协程,其他编程语言的开发机制是一般基于线程的,开启过多线程，
资源消耗非常大,这里就突显Golang在并发上的优势了.
*/

/*
MPG模型基本介绍
1.M:操作系统的主线程(是物理线程)
2.P:协程执行需要的上下文
3.G:协程
*/

// 编写一个函数,每隔1s输出"hello world"
func test1() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("test1() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func test2() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("test2() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func test3() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("test3() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test1() //开启协程 goroutine,test()和main()同时执行,以主线程为主,完成后整个函数结束,协程没有完成也会退出
	go test2() //开启协程
	go test3() //开启协程
	for i := 1; i <= 1000; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
