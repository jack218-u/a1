package main
import "fmt"
func main(){
	a := 4
	squareVal(a)
	squareAdd(&a)
	fmt.Println(&a)
}
func squareVal(v int){
	v *=v
	fmt.Println(&v,v)
}
func squareAdd(p *int){
	*p *= *p
	fmt.Println(p,*p)
}