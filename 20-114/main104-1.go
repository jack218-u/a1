package main
import "fmt"
func main(){
	//here:
	for i :=0; i < 2; i++ {
		for j :=1; j < 4; j++ {
			if j == 2 {
				// break和下面的continue+循环标签功能一样
				continue  //跳出本层循环进入下一个循环
			}
		fmt.Println("i=",i,"j=",j)	
		}
	}
}