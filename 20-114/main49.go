package main
import (
	"fmt"
	_ "unsafe"
)
func main(){
	var num1 int = 100
	var num2 float32 = 23.254
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%T",b)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("str type %T str=%q\n",str,str)

}