package main
import "fmt"
func main(){
	var char byte  
	fmt.Println("输入小写字母a-z")
	fmt.Scanf("%c",&char)
	switch char {
	case 'a' :
		fmt.Println("A")
	case 'b' :
		fmt.Println("B")
	case 'c' :
		fmt.Println("C")
	case 'd' :
		fmt.Println("D")
	case 'e' :
		fmt.Println("E")
	 default: 
		fmt.Println("others")
	}
}