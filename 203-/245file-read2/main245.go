// 读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中),这种方式适用于文件不大的情况
// 相关方法和函数(ioutil.ReadFile) //此函数已经弃用
package main

import (
	"fmt"
	"os"
)

func main() {
	//使用ioutil.ReadFile函数读取文件
	file := "d:/practice/245.txt"
	// content, err := ioutil.ReadFile(file) //此函数已经被弃用,用os.ReadFile函数代替
	content, err := os.ReadFile(file) //隐式输出
	if err != nil {
		fmt.Println("read file err=", err)
	}
	fmt.Printf("%s", content) //[]byte切片所有需要%s参数显示，类似转成string
	fmt.Printf("%T", content)
	//把读取的内容显示到终端
	//因为我们没有显式的Open文件,因此也不需要显式的Close文件
	//因为,文件的Open和Close被封装到ReadFile函数内部
}
