package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}() //匿名函数结尾要加(),如果有参数还要在结尾的()传入参数
	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

/*
goroutine中使用recover,解决协程中出现panic导致程序崩溃问题处理比如一个协程,
但是这个协程中出现了panic,导致整个程序崩溃,我们希望这个panic不导致整个程序崩溃,
而是继续执行,可以在goroutine中使用recover来捕获panic进行处理,这样即使这个协程发生问题,
但是主线程仍然不受影响,可以继续执行
*/
func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main()ok=", i)
		time.Sleep(time.Second)
	}
}
