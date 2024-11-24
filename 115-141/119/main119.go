package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}
func myFun(funvar func(int, int) int, num1 int, num2 int) int { //函数可以当成是形参调用
	return funvar(num1, num2)
}

//对函数类型做自定义
type myFunType func(int, int) int

//自定义函数数据类型以后 引用自定义以后的函数数据类型
func myFun1(funvar myFunType, num1 int, num2 int) int { //函数可以当成是形参调用
	return funvar(num1, num2)
}

//支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

//可变参数,args这里表示数组切片
func sum(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 1; i < len(args); i++ { //len(args)表示args的长度
		sum += args[i] //args[i]表示取切片的第n个元素值，从0开始
	}
	return sum
}
func sum2(n1 int, args ...int) int {
	var sum2 int = n1
	//遍历args,用for range方式实现
	for _, value := range args { //key,value表示数组中的序号和数值，因为不需要返回序号，所以用_占用省略
		sum2 += value //value表示args中第n个元素值，从0开始
	}
	return sum2
}

func main() {
	a := getSum
	fmt.Printf("a的数据类型%T \ngetsum的数据类型%T\n", a, getSum)
	res := a(10, 40)
	fmt.Println("res的值=", res)
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2的值=", res2)
	type myInt int //自定义一个数据类型，类型为int，重新定义之后，数据类型变成不同的数据类型，和int数据类型赋值需要做显式转换，
	var num1 myInt = 10
	var num2 int = int(num1) //这里需要显式转换
	fmt.Printf("num1的数据类型是%T,num1的值=%v\nnum2的数据类型是%T,num2的值=%v\n", num1, num1, num2, num2)
	res3 := myFun1(getSum, 500, 600)
	fmt.Printf("res3的值=%v\n", res3)
	res4, res5 := getSumAndSub(800, 600)
	fmt.Printf("res4的值=%v,res5的值=%v\n", res4, res5)
	//测试可变参数的使用
	res6 := sum(10, 20, 70, -50)
	fmt.Printf("res6的值=%v\n", res6)
	res7 := sum2(10, 20, -30, 168)
	fmt.Printf("res7的值=%v", res7)

}
