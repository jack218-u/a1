package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "d:/practice/test2.txt" //定义需要调用变量以及文件路径,变量名为file
	//在原来基础上末尾追加写入
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666) //调用文件file，返回新的 变量f
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer f.Close() //这里需要写变量f
	str := "ABC-ENGLISH\n"
	writer := bufio.NewWriter(f) //调用变量f
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
