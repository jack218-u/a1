package main
import "fmt"
func main(){
	var n = 6
	for i := 1; i <= 6; i++ {
		fmt.Printf("%v + %v = %v\n",i, n-i,n)
	}
}