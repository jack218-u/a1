package main
import "fmt"
//for - range遍历
func main(){
	var str string = "hellogolang"
	for i := 0; i < len(str);i++ {       //i是字符串的下标，从0开始，len(str)是字符串的长度，所以这里的i 需要小于 len(str),不能等于
		fmt.Printf("%c,%d\n",str[i],i)
	}
}