package main

import "fmt"

//ch chan<- int,这样ch就只能写操作
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

//ch <-chan int，这样ch就只能读操作
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	//管道可以声明为制度或者只写
	//1. 在默认情况下,管道是双向的
	//var chan1 chan int //可读可写

	//2.声明为只写
	chan2 := make(chan<- int, 3) //只写的channel
	chan2 <- 20
	// num := <-chan2 //只写的channel 不能从它读取数据,否则报错
	fmt.Println(chan2)

	//3.声明为只读
	chan3 := make(<-chan int)
	fmt.Println(chan3)

	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)
	var total = 0
	for range exitChan {
		total++
		if total == 2 {
			//如果total ==这里写3，就会堵塞，因为exitChan的容量是2,而且两个协程send和
			//recv执行完输入一个空的结构体到exitChan管道,最多exitChan中只能有两个数据
			break
		}
	}
	fmt.Println("结束")
}
