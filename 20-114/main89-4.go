package main
import "fmt"
func main(){
	var week string
	fmt.Println("输入星期日期")
	fmt.Scanln(&week)
	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("隆江猪脚饭")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	 default: 
		fmt.Println("大盘鸡")
	}
}