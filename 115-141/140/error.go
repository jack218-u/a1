package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型
// 1. error.New("错误说明")，会返回一个error类型的值，表示一个错误
// 2. panic内置函数，接收一个interface{}类型的值(也就是任何值了)作为参数。可以接收error类型的变量，输出错误信息并退出程序

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config2.ini") //调用上面的readConf函数这里写config2和上面的config不一致，会出现报错信息，如果一致就会继续执行
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}
func main() {
	test02() //调用test02()函数
	fmt.Println("main()下面的代码...")
}
