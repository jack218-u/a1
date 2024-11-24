package main

import "fmt"

func peach(n int) int { //注意函数返回类型为int，如果换成是uint8，则会出现数据溢出，结果不正确，uint8的范围是0--255
	//猴子吃桃子问题，第一天吃总桃子数量一半，再多吃一个，以后每天如此，到第十天的时候，发现只有一个，求第一天桃子原本数量
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
func main() {
	fmt.Println(peach(1))
	fmt.Println("peach(2)=", peach(2))
	fmt.Println("peach(3)=", peach(3))
	fmt.Println("peach(4)=", peach(4))
	fmt.Println("peach(5)=", peach(5))
	fmt.Println("peach(6)=", peach(6))
	fmt.Println("peach(7)=", peach(7))
	fmt.Println("peach(8)=", peach(8))
	fmt.Println("peach(9)=", peach(9))
}
