package main
import "fmt"
func main(){
	label2:
	for i :=0 ;i < 4 ; i++ {
	// label1:
		for j :=0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("i=",i,"j=",j)
		}
	}
}