package main
import "fmt"
func test(n1 int){
	n1 = n1 + 1
	fmt.Println("test() n1=",n1) //?输出结果
}
//一个函数getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ",n1)  //30
	//当函数有return语句时，就是将结果返回给调用者，即谁调用我，我就返回给谁
	return sum //此处需要有return返回值，因为main主函数中有sum := getSum(10,20),否则语法错误
}
//请编写函数，可以计算两个数的和与差
func sumAndsub(n1 int,n2 int) (int,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum,sub
}

func main(){
	n1 := 10
	//调用test
	test(n1)
	fmt.Println("main() n1=",n1)

	sum := getSum(10,20)  //需要getSum函数中有return
	fmt.Println("main sum = ",sum) //30
	
	//调用sumAndsub
	res1,res2 := sumAndsub(1,2)
	fmt.Printf("res1=%v, res2=%v\n",res1,res2)
	//希望忽略某个值，则用_号表示占位符
	_,res3 := sumAndsub(3,5)
	fmt.Printf("res3=%v",res3)
}