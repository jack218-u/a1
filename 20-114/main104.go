package main
import "fmt"
func main(){
	for i :=0; i < 3; i++ {
		for j :=0; j < 4; j++ {
			if j == 2 {
				// break
				continue //结束当前j,继续进行下一次迭代
			}
		fmt.Println("i=",i,"j=",j)	
		}
	}
}