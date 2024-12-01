package main

import (
	"fmt"
)

func main() {
	//使用select解决从管道取数据堵塞问题
	//1. 定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统方法遍历管道时,如果不关闭会阻塞而导致deadlock
	//问题,在实际开发中,不好确定什么时候关闭管道
label:
	for {
		select {
		//注意:如果管道最终一直没有关闭,不会一直阻塞而死锁(deadlock)
		//会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			// time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			// time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了,不玩了,可以加入自定义逻辑\n")
			// time.Sleep(time.Second)
			// return
			// 此处如果用return则整个for循环关闭,return后面步骤的代码不能执行
			break label
		}
	}
	fmt.Println("abc")
}
