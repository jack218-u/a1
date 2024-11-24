package main
import "fmt"

func main(){
	//有两个变量a和b，要求对其交换，但是不能使用中间变量，打印输出结果
	var a int = 20
	var b int = 10
	// a = a + b
	// b = a - b //b = (a + b) - b ==> b = a ,这里b换成了a(原始的a)
	// a = a - b //a = (a + b) - a(备注,本来这里应该写b，但此时b已经替换成a了，参考上一行) = b (原始的b)
	fmt.Println(a,b)  //转换a和b
}


