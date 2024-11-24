package main
import "fmt"
func main() {
	//如果return是在普通的函数，则表示跳出该函数，即不再执行函数中return后面的代码，也可以理解成终止函数
	//如果return是在main函数，表示终止main函数，也就是终止程序
	for i :=1; i <=3; i++ {
		for j :=1; j <=4; j++ {
			if j == 3 {
				// continue //跳过2，后面不输出j=2，继续后面的迭代执行，继续该层循环的j=3
				// break   //遇到2，停止该层循环执行，不输出后面的迭代j=2，以及后面的j=3，跳到下一层i的循环
				return     //停止整个函数或者方法，只输出j == 2前面的部分，停止该层循环，以及外层i层的循环
			}
			fmt.Printf("i=%v,j=%v \n",i,j)
		}
	}
}