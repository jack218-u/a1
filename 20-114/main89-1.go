package main
import "fmt"
import "strings" //小写字母转换为大写字母需要用strings函数
func main(){
	// var char byte  //如果定义为byte字符类型，后面则不能用strings函数转换
	var char string  //这里需要用到strings函数转换，此处数据只能定义为string字符串类型
	fmt.Println("输入小写字母a-z")
	fmt.Scanln(&char)
	switch char {
	case "a","b","c","d","e" :  //字符串需要用英文双引号圈起来
	// case 'a','b','c','d','e' :  //如果是用单引号圈起来，不符合string类型的语法规则，单引号引用的是常量
		fmt.Println(strings.ToUpper(char))
	 default: 
		fmt.Println("others")
	}
}