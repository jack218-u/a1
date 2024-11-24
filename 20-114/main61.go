package main
import "fmt"
func main(){
	//假如你还有有97天放假，问：等于几个星期零几天
	var days int = 97
	var weeks int = days / 7
	var leftdays = days % 7
	fmt.Printf("%d个星期零%d天\n",weeks,leftdays)
	//定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为:
	//5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100) //此处5需要加小数点，如果是5/9,按照计算机计算则是0
	fmt.Printf("%v对应的摄氏温度%v",huashi,sheshi)
}