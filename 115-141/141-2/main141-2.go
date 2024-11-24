package main

import (
	"fmt"
	"math/rand"
)

func randnum() int {
	n := rand.Intn(10) + 1
	return n
}
func main() {
	f := randnum()
	for i := 1; i <= 11; i++ {
		var count int
		var num int
		fmt.Println("请输入您认为的随机生成的数字")
		fmt.Scanln(&num)
		count += i
		if num == f {
			if count == 1 {
				fmt.Println("你真是个天才")
			} else if count == 2 || count == 3 {
				fmt.Println("你很聪明")
			} else if count >= 4 && count <= 9 {
				fmt.Println("一般般")
			} else if count == 10 {
				fmt.Println("你可算猜对了")
			} else {
				fmt.Println("说你点啥好呢")
			}
			break
		}
	}
}
