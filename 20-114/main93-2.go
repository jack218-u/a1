package main
import "fmt"
//for - range遍历
func main(){
	var str string = "hellogolang你好"
	for key,value := range str {  //for range是按照字符方式遍历的，如果字符串有汉字也是ok，按照汉字占用字节输出字符所在的序号
		fmt.Printf("%d,%c\n",key,value)
	}
}