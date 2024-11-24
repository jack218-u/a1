package main
import (
	"fmt"
	"math"
)
func main(){
	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0
	m := b * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Println(x1,x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Println(x1)
	} else if m < 0 {
		fmt.Println("无解")
	}
}