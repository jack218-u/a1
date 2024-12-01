package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <= 800000; num++ {
		flag := true
		//判断num是不是素数
		for i := 2; i < num; i++ { //此处i不能=num,否则此情况下num%i == 0,不符合预期
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// //将这个数放入到primeChan
			// primeChan <- num
		}
	}
	end := time.Now().Unix()
	fmt.Println("使用普通方法耗时=", end-start)
}
