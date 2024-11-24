package main
import "fmt"
func main(){
	var month byte
	fmt.Println("月份")
	fmt.Scanln(&month)
	switch month {
	case 3,4,5:
		fmt.Println("春季")
	case 6,7,8:
		fmt.Println("夏季")
	case 9,10,11 :
		fmt.Println("秋季")
	 default: 
		fmt.Println("冬季")
	}
}