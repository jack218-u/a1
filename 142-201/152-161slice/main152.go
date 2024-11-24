package main

import "fmt"

func main() {
	//1. 切片的英文是slice
	//2. 切片是数组的一个引用，因此切片是引用类型，在进行传递的时候，遵守引用传递
	// 从底层看，其实就是一个数据结构(struct结构体)
	//3. 切片的使用和数组类型，遍历切片、访问切片的元素和求切片长度len(slice)都一样
	//4. 切片的长度可以变化，因此切片是一个可以动态变化的数组
	//5. 切片定义的基本语法：var 变量名 []类型 比如var a []int  []中不需要写长度，和数组不同
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//1. 声明/定义一个切片
	//2. intArr[1:3] 表示 slice引用到intArr这个数组
	//3. 引用intArr数组的起始下标为1，终止下标为3(但是不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice 的元素个数是=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) //切片的容量是可以动态变化
	fmt.Printf("intArr[1]的地址%p intArr[1]=%v\n", &intArr[1], intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0]=%v\n", &slice[0], slice[0]) //slice[0]和intArr[1]值一样
	fmt.Println()
	//4.对切片的修改会修改被引用的数组中元素的值
	slice[1] = 34 //修改切片slice[1]是intArr[2],所以intArr数组里面下标为2的元素变成了34
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)
	println()
	intArr[1] = 23 //修改切片slice[1]是intArr[2],所以intArr数组里面下标为2的元素变成了34
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println()
	//切片使用的三种方式
	//1. 定义一个切片，引用一个已经创建好的数组
	//2. 通过一个内置函数make创建一个切片, var slice []int = make([]int,4,10) //make的参数分别为类型，长度，容量
	//var 切片名 []type = make([]type,len,[cap]) //type是数据类型，len 长度 cap 切片容量，为可选项
	//容量必须大于等于长度，len >= cap
	//切片的使用make,通过make方式创建的切片对应的数组由make底层维护，对外不可见，即只能通过slice去访问各个元素
	//如果没有给切片各个元素赋值，那么切片就会使用默认值，int float类型=>0,string=>"",bool=>false
	var slice2 []float64 = make([]float64, 5, 10)
	//对于切片，必须make使用
	//通过make方式创建
	fmt.Println(slice2)
	slice2[1] = 10
	slice2[3] = 30
	fmt.Println(slice2)
	fmt.Println("slice2的size=", len(slice2))
	fmt.Println("slice2的cap=", cap(slice2))
	fmt.Println()
	//4. 定义一个切片，直接就指定具体数组，使用原理类似make方式,如果不指定，默认为空即 var slice []int
	// 使用第三种方式创建切片
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice))
	fmt.Println("strSlice cap=", cap(strSlice))
	var slice3 []float64
	fmt.Println("slice3=", slice3) //没有赋值，为空不能使用
	var slice4 []float64 = []float64{1, 2, 3, 4, 5, 6}
	fmt.Println(slice4)
	//总结：
	//方式1是直接引用数组，这个数组是事先存在的，可见的
	//方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，不可见
}
