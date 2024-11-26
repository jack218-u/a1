package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包:synchornized同步
	//Mutex:是互斥
	lock sync.Mutex
)

// test 函数就是计算n!,n的阶乘,让这个结果放入到myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res //concurrent map writes 不加锁会出现这个问题
	//解锁
	lock.Unlock()
}
func main() {
	for i := 1; i <= 10; i++ { //i不能设置过大,否则阶乘的结果超过int限制
		go test(i)
	}
	//休眠10秒钟,【第二个问题】
	time.Sleep(time.Second * 5) //等待5秒,预期5秒内阶乘能完成计算,否则执行主程序会让协程退出
	//这里我们输出结果,变量这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
