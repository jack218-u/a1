package main
import "fmt"
func main(){
	// 	某人有100000现金，每经过一次路口需要交费，规则如下：
	// 1.当现金大于50000时，每次交费5%；
	// 2.当先进小于等于50000时，每次交1000；
	// 编程计算该人经过多少次路口，使用go编程语言的for break功能实现
	var cash float64 = 100000
	var count int = 0
	for {
		if cash > 5000 {
			cash *= 0.95
		} else if cash <=5000 && cash >= 1000 {
			cash -= 1000
		} else {
			break
		}
		count++
		fmt.Println(count)
	} 
}