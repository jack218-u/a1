package main
import "fmt"
func main(){
	//输出1到100能被9整除的数的个数和总和
	var count int = 0
	var sum int = 0
	for i :=1 ;i <= 100;i++{  //i从1开始规避的0
		if i % 9 == 0 {       //能被9整除，即取余为0
			count++
			sum +=i
		}
	}
	fmt.Println(count,sum)
}