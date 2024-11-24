package main
import "fmt"
func main(){
	//九九乘法表
	var Num int 
	fmt.Println("输入想实现的乘法表的数字")
	fmt.Scanln(&Num)
	for i :=1; i <= Num; i++ {
		for j :=1; j <=i; j++ {
			fmt.Printf("%v*%v=%v\t",j,i,j*i)
		}
		fmt.Println()
	}
}