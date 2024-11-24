package main
import "fmt"
func main(){
	var user string 
	var passwd string
	for i := 1; i <=3 ; i ++ {
		fmt.Println("输入用户名")
		fmt.Scanln(&user)
		fmt.Println("输入密码")
		fmt.Scanln(&passwd)
		if user == "张无忌" && passwd == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("还有%v次机会\n",3-i)
		}
	}
}