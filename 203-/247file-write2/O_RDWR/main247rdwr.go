package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "d:/practice/test3.txt"                          //定义需要调用变量以及文件路径,变量名为file
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666) //调用文件file，返回新的 变量f
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer f.Close() //这里需要写变量f
	//先读取原来的文件内容,并显示在终端
	reader := bufio.NewReader(f)
	for {
		a, err := reader.ReadString('\n') //读取一行的内容,直到换行为止
		if err == io.EOF {                //如果读取到文件末尾，就结束循环
			break
		}
		fmt.Print(a) //显示到终端,因为读的时候就是按行读取，所以这里也按行显示,不加换行符ln
	}
	//读完之后执行写入操作
	str := "hello,北京\n"
	writer := bufio.NewWriter(f) //调用变量f
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
