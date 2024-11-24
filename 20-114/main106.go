package main
import "fmt"
func main(){
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num ==0 {
			break
		} else if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("输出正数的个数%v,输出负数的个数%v",positiveCount,negativeCount)
}