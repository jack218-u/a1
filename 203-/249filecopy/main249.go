package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
将一个图片/电影/音乐文件拷贝到另一个文件e:/abc.jpg io包
func Copy(dst Writer,src Reader)(written int64,err error)
注意Copy函数是io包提供的
*/
//编写一个函数,接收两个文件路径
func copy(dst, src string) (written int64, err error) {
	a, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer a.Close()
	//通过a获取Reader
	reader := bufio.NewReader(a) //显式输出用os.buffio.NewReader()函数,隐式输出用os.ReadFile()函数
	//打开dst文件
	c, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	//通过c获取到Writer
	writer := bufio.NewWriter(c)
	defer c.Close()
	return io.Copy(writer, reader)
}
func main() {
	src1 := "d:/practice/1/1.png"
	dst1 := "d:/practice/2/2.png"
	a, err := copy(dst1, src1)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 err=%v", err)
	}
	fmt.Println(a) //为拷贝字节数
	os.Open(src1)
	b, err := os.Open(src1)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	fmt.Println(*b) //os.Open(src1)函数中返回的文件为指针类型,需要用*解引用
}
