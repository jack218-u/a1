package main
import "fmt"

func main(){
	var n1 int = 10
	var n2 int = 41
	n3 := n1 + n2
	if n3 > 50 {
		fmt.Println("50")
	}
	var n4 float32 = 11.20
	var n5 float32 = 12.00
	if n4 > 11 && n5 < 20{
		fmt.Println("和= ",n5+n4)
	}
    var num1 int32 = 10
    var num2 int32 = 5
	if (num1 + num2) % 5 == 0 && (num1 + num2 ) %3 == 0 {
		fmt.Println("能被3整除又能被5整除")
	}
	var year int =2020
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println("闰年")
	}
}
