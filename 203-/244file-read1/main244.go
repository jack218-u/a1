package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//概念说明：file的叫法
	//1. file叫file对象
	//2. file叫file指针
	//3. file叫file文件句柄
	file, err := os.Open("d:/practice/244.txt") //Open函数参数比OpenFile少，用于只读的操作
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//当函数退出时,要及时的关闭file
	defer file.Close() //要及时的关闭file句柄,否则有文件内存泄漏
	//创建一个 *Reader, 是带缓冲的
	//默认缓冲区大小是4096
	/*
		const (
			defaultBufSize = 4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		// \n读取一行的内容,直到换行为止
		str, err := reader.ReadString('\n') //调用bufio.NewReader中的ReadString()方法
		//读到一个换行符就结束
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束") //如果显示的"文件读取结束"上还有空格，说明文件末尾还有空行
}
