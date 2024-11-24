//写文件操作应用实例-基本介绍
/*
func OpenFile(name string,flag int,perm FileMode) (file *File,err error)
flag 文件打开的方式,比如只读,只读,只写,读写等,其中O_TRUNC表示清空,需要特别注意，否则导致数据丢失,可以用|符号组合使用
perm 权限控制,在Windows下无效
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件,写入内容5句"hello ,Gardon"
	//1. 打开文件 d:/abc.txt
	filePath := "d:/def.txt"                                          //需要调用文件的路径
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //file是变量
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入5句话"hello,Gardon"
	str := "hello,Gardon\n" // \n表示换行
	//写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file) //写入19行的变量
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为Writer是带缓存的,因此在调用WriteString()方法时,
	//其实内容是先写入到缓存,所以需要调用Flush方法,将缓存的数据真正写入到文件中,否则文件中会没有数据
	writer.Flush()
}
