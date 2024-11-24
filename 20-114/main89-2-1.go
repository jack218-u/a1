package main
import "fmt"
func main(){
	var score float64
	fmt.Println("输入学生成绩")
	fmt.Scanln(&score)
	switch int(score / 60){
	case 1 :
		fmt.Println("合格")
	case 0 :
		fmt.Println("不合格")
	 default: 
		fmt.Println("输入有误")
	}
	fmt.Println(score / 60)
}