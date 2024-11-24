package main

import (
	"fmt"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //此条件表示err为空，证明 文件或者目录存在---> 路径存在的判断
		return true, nil //返回true，表示文件存在，同时返回nil表示err为空---> 路径存在的输出
	}
	if os.IsNotExist(err) { //说明此函数判断一个文件或者目录不存在 为真,在这里表示err不存在为真---> 路径不存在的判断
		return false, nil //返回false，表示文件不存在，同时返回nil表示err为空---> 路径不存在的输出
	}
	return false, err //---> 其他类型的错误判断和错误输出,比如这些错误可能是权限问题、磁盘错误等。在这种情况下，路径的状态是不确定的
}
func main() {
	//将d:/practice/1/abc.txt文件导入到2/kkk.txt
	//1. 首先将abc.txt内容读取到内存
	//2. 将读取到的内容写入到kkk.txt
	file1 := "d:/practice/1/abc.txt"
	file2 := "d:/practice/2/kkk.txt"
	data, err := os.ReadFile(file1) //隐式的打开和关闭文件
	if err != nil {
		fmt.Printf("read file err= %v\n", err)
		return
	}
	//读取的数据data写到了新的文件file2中，全部覆盖写，file2原先的文件被清空，然后写入data文件中的数据
	err = os.WriteFile(file2, data, 0666) //隐式的打开和关闭文件,将上面的err重新利用
	if err != nil {
		fmt.Printf("write file err= %v\n", err)
	}
	var a bool
	var b error
	file3 := "d:/practice/1/abc2.txt"
	a, b = PathExist(file3)
	fmt.Println(a, b)
}
