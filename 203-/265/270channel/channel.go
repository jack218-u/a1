/*  为什么需要channel
通过使用全局变量加锁同步来解决goroutine并发问题,但不完美
1.主线程等待所有gotoutine全部完成的时间很难确定,我们这里设置10s,仅仅是估算
2.如果主线程休眠时间长了,会加长等待时间,如果等待时间短了,
可能还有goroutine处于工作状态,这时也会随主线程的退出而销毁
3.通过全局变量加锁同步来实现通讯,也并不利用多个协程对全局变量的读写操作.
4.上面的分析都在呼唤一个新的通讯机制:管道
*/
/*
   channel介绍
   1.channel本质就是一个数据结构-队列
   2.数据是先进先出【FIFO】
   3.数据安全,多goroutine访问时,不需要加锁,就是说channel本身就是线程安全的
   即:多个协程操作同一个管道时,不会发生资源竞争问题
   4.channel是有类型的,比如一个string类型的channel只能存放string类型数据,
   如果想混用,把管道声明成空接口
*/
/*
定义 / 声明 / channel
举例:
var intChan chan int(intChan用于存放int数据)
var mapChan chan map[int]string(mapChan用于存放map[int]string数据)
var interfaceChan chan interface{}(interfaceChan用于存放任意类型数据)
var perChan chan Person //Person结构体的管道,Person需要先定义
var perChan2 chan *Person //存放Person结构体指针,Person需要先定义
说明:
1.管道是引用类型
2.管道必须初始化才能写入数据,即make后才能使用
3.管道是有类型的,只能写入指定类型的数据,比如intChan只能写入int数据,interfaceChan可以写入任意类型数据
4.管道是先进先出,所以不能用range遍历管道,range遍历管道会阻塞协程,直到管道关闭
5.管道是协程安全的,多个协程可以同时操作一个管道,不会发生资源竞争问题
6.管道是协程安全的,但是不能用于多个协程之间传递指针,因为指针是可变的,多个协程操作同一个指针,可能会导致数据混乱
*/
/*
channel使用注意事项:
1.channel中只能存放指定的数据类型
2.channel的数据放满后,就不能再放入了
3.如果从channel中取出数据后,可以继续放入
4.在没有使用协程的情况下,如果channel数据取完了,再取就会报dead lock
*/
package main

import (
	"fmt"
)

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	intChan := make(chan int, 100)
	//2. 看看intChan是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//4.看看管道的长度和cap容量,管道和map不一样不能自动扩容
	//长度2因为加了两个数据,容量3,因为make的时候设置的是3
	//注意:给管道写入数据,不能超过其容量
	intChan <- 50
	intChan <- 60
	intChan <- 70
	// intChan <- 100
	fmt.Printf("channel len=%d cap=%d\n", len(intChan), cap(intChan))
	//5.从管道中读取数据
	num2 := <-intChan // 从intChan通道接收一个整数值并将其赋给num2变量
	<-intChan         //取值以后不接收也可以,没有接收的变量
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%d cap=%d\n", len(intChan), cap(intChan))
	//6.在没有使用协程的情况下,如果我们的管道数据已经全部取出了,再取就会报告deadlocak
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)
	fmt.Printf("channel len=%d cap=%d\n", len(intChan), cap(intChan))
	// num5 := <-intChan
	// fmt.Println("num5=", num5)
	//7.关闭管道
	close(intChan)
}
