package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "d:/practice/test1.txt"                           //定义需要调用变量以及文件路径,变量名为file
	f, err := os.OpenFile(file, os.O_TRUNC|os.O_WRONLY, 0666) //调用文件file，返回新的 变量f
	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer f.Close() //这里需要写变量f
	str := "hello,world\n"
	writer := bufio.NewWriter(f) //调用变量f
	for i := 0; i < 11; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
