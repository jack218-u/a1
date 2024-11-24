package main
import "fmt"

func main(){

	var i int = 10
	i += 17
	fmt.Println(i)

	a,b := 2,9
	a,b = b,a
	fmt.Println(a,b)
}