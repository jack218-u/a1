package main
import "fmt"

func main(){

	var score int
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("BMW")
	} else if score >= 80 && score <= 90 {
		fmt.Println("ipad")
	} else {
		fmt.Println("no")
	}
}
