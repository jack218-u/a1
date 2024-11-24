package main

import (
	"fmt"
)

func main() {
	var year int
	var month int
	var day int
	for i := 1; ; i++ {
		fmt.Print("请输入年:") //注意此处如果是Println，则输如框在下一行，而不是在同一行
		fmt.Scanln(&year)
		fmt.Print("请输入月:")
		fmt.Scanln(&month)
		if month <= 0 {
			fmt.Println("月份不正确")
			continue //结束本次循环开始下次循环，从输入年份开始
		}
		fmt.Print("请输入日:")
		fmt.Scanln(&day)
		fmt.Printf("您输入的日期是: %v年%v月%v日\n", year, month, day)
	}

}
