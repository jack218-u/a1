package main
import "fmt"
func main(){
	//嵌套分支
	var month byte
	var age byte
	var price float64 = 60.00
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age >= 60 {
			fmt.Printf("%v月,%v岁,%v元",month,age,price/3)
		} else if age >= 18 {
			fmt.Printf("%v月,%v岁,%v元",month,age,price)
		} else {
			fmt.Printf("%v月,%v岁,%v元",month,age,price/2)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人票价40")
		} else {
			fmt.Println("淡季儿童和老人票价20")
		}
	}
	
}