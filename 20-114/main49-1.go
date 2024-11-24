package main
import (
	"fmt"
	"strconv"
)
func main(){
	var num3 int = 99
	var num4 float32 = 23.254
	// var b2 bool = true
	var str string //空的str
	str = strconv.FormatInt(int64(num3),10) //FormatInt默认是int64进制的，需要转，10表示转换成10进制的
	fmt.Printf("str type %T str=%q\n",str,str)
	str = strconv.FormatFloat(int64(num4),'f',10,10) 
	fmt.Printf("str type %T str=%q\n",str,str)

}