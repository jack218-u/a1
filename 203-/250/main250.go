package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体用于保存统计的结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字的个数
	SpaceCount int //记录空格的个数
	OtherCount int //记录其他字符的个数
}

func main() {
	//思路:打开一个文件,创一个Reader
	//每读一行,就去统计该行有多少个 英文 数字 空格和其他字符
	//然后将结果保存到一个结构体中
	fileName := "d:/practice/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	//定义一个CharCount结构体实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)
	//开始循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n') //每次只读一行读到\n结束,然后循环读下一行
		if err == io.EOF {                  //读到文件末尾退出
			break
		}
		//遍历str,进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z': //对字符比较，会比较他们的ascii码值
				fallthrough //穿透处理,继续叠加数字
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
		fmt.Println(str)
	}
	//输出统计结果看看是否正确
	fmt.Printf("英文个数:%d\n数字个数:%d\n空格个数:%d\n其他字符个数:%d\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
