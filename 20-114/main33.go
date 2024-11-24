package main
import "fmt"
import "unsafe"
var (
	a1 int = 10
	a2 int = 11
	a3 int = 12
)
func main(){
	var i int 
	fmt.Println("i=",i)   //不定义变量i，默认值为0
	var num = 10.11
	fmt.Printf("num的类型为:%T,num的值为:%v\n",num,num)
	num2 := 10.11
	fmt.Printf("num2的类型为:%T,num2的值为:%v\n",num2,num2)
	var num3 int = 10
	fmt.Printf("num3的类型为:%T,num3的值为:%v\n",num3,num3)
	var n1, n2, n3 int 
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	fmt.Println(a1,a2,a3)
	n4 := 10
	fmt.Printf("n1占用的字符数: %d\n",unsafe.Sizeof(n4))
	var n5 float32 = 123456789.1234567890123456789
	var n6 float64 = 123456789.1234567890123456789
	fmt.Println(n5,n6)
}