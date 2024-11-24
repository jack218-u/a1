package main
import "fmt"
func main(){
	//1+2+3+4...+100,求第一次大于20时候当前的数
	sum :=0
	for i := 1; i <= 100; i++ {
		sum += i 
		if sum >= 20 {
			fmt.Println(i)
			break
		}
	} 
}