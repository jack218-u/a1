package main

import "fmt"

func main() {
	//希望一次性注释，选择需要注释的部分，ctrl + / ，重新执行表示 取消注释
	// /*   */块注释，块注释里面不能再有块注释
	// 规范的代码风格：尽量用行注释
	// fmt.Println("hellogolang")
	// fmt.Println("hellogolang")
	/*
				fmt.Println("hellogolang")
		fmt.Println("hellogolang")
				fmt.Println("hellogolang")
				fmt.Println("hellogolang")
	*/
	fmt.Println("hellogolang")
	fmt.Println("hellogolang")
	fmt.Println("hellogolang")
	fmt.Println("hellogolang") //在执行框中输入gofmt 代码名表示自动缩进，不报错，加 -w参数自动报错，自动缩进格式
	fmt.Println("tom\tjack")   //表示一个制表符，通常可以用它排版
	fmt.Println("hello\ngolang")
	fmt.Println("hello\\golang")
	fmt.Println("hello\"golang\"")
	fmt.Println("天龙八部雪山飞狐\r张飞") // \r从当前最前面开始输出覆盖掉前面的内容
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12 \t河北\t北京")
}
