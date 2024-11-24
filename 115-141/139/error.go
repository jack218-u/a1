package main

import (
	"fmt"
	"time"
)

// 1.默认情况下，当发生错误后(panic),程序就会退出(崩溃)
// 2.如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在捕获到错误后，给管理员一个提示(邮件，短信)
// 3.引出错误处理机制
// go错误处理机制：defer，panic，recover，go抛出一个panic的异常，
// 然后在defer中通过recover捕获这个异常，然后正常处理,让代码更健壮
func test() {
	//使用defer和recover来捕获和处理异常
	defer func() {
		err := recover() //recover()为内置函数，可以捕获到异常
		if err != nil {  //说明捕获到异常
			fmt.Println("err=", err)
			//这里就可以将错误信息发送给管理员
			fmt.Println("发送邮件给admin@sohu.com")
		}
	}() //()表示调用匿名函数
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
func main() {
	//测试
	test()
	for {
		fmt.Println("main()下面的代码...")
		time.Sleep(time.Second)
	}
}
