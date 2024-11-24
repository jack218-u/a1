package main
import "fmt"
func main(){
	//逻辑运算符：true和false
	var age int = 40
	//逻辑与&&，两个都成立才成立，输出结果,也称短路与，第一个条件为false，结果为false，后面的不会判断
	if age > 30 && age < 40 {
		fmt.Println("OK1")
	}
	if age > 30 && age < 50 {
		fmt.Println("OK2")
	}
	//逻辑或||，只要有一个成立即成立，输出结果，也称短路或，第一个条件为true，结果为true，后面的不会判断
	if age > 30 || age < 40 {
		fmt.Println("OK3")
	}
	if age > 30 || age < 50 {
		fmt.Println("OK4")
	}
	if age < 40 || age > 30 {
		fmt.Println("OK5")
	}
	//逻辑非!,如果比较结果是true，则输出false，如果结果是false，则输出true
	if !(age < 30){
		fmt.Println("OK6") //age < 30 ，其中age = 40 是false，则此语句逻辑非是true，能输出OK6
	}
	if !(age > 60){
		fmt.Println("OK7") //age < 60 ，其中age = 40 是false，则此语句逻辑非是true，能输出OK7
	}
	if !(age > 30){
		fmt.Println("OK8") //age < 30 ，其中age = 40 是true，则此语句逻辑非是false，不能输出OK8
	}
}