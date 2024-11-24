package main
import "fmt"
func main(){
	var n1 int32 = 12
	var n4 int8
	n4 = int8(n1) + 127
	fmt.Println(n4)
}