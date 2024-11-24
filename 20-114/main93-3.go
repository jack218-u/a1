package main
import "fmt"
//for - range遍历
func main(){
	var str string = "hellogolang你好" //因为有汉字，需要转换成数组才能按照序号输出
	str2 := []rune(str)  //把str转换成[]rune类型
	for i := 0; i < len(str2);i++ {       //i是字符串的下标，从0开始，len(str)是字符串的长度，所以这里的i 需要小于 len(str),不能等于
		fmt.Printf("%c,%d\n",str2[i],i)
	}
}