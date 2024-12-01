/*
常量基本介绍
1.常量使用const修改
2.常量在定义的时候,必须初始化
3.常量不能修改
4.常量只能修饰bool、数值类型(int,float系列)、string类型
5.语法:const identifier [type]= Value
6.Golang中没有说明常量名必须字母大写的规范
7.仍然通过首字母大小写来控制常量的访问范围,首字母大小可以在其他包中访问,小写只能在本包访问

在 Go 语言中，常量（以及变量、函数等）的可见性（即访问范围）
可以通过其名称的首字母大小写来控制。具体规则如下：

首字母大写：表示该常量是 导出的（exported），可以在其他包中访问。
首字母小写：表示该常量是 未导出的（unexported），只能在当前包内访问。
*/
package main

import "fmt"

func main() {
	var num int
	//常量声明的时候,必须赋值
	const tax int = 0
	// tax = 10 //常量是不能修改
	fmt.Println(num, tax)
	const name = "tom" //常量声明也可以用类型推导方式,不写数据类型
	fmt.Println(name)
	// const a int //常量定义的时候必须初始化,赋值
	const b = 9 / 3
	fmt.Println(b) //固定的值可以
	// const c = num / 3 //错误写法 因为num是变量 const是常量
	// const c = getVal() //getVal返回值可变
	// const()写在main函数里面就是局部常量,写在main外面就是全局常量
	const (
		a = 1
		c = 2
	)
	//常量以小括号方式声明
	fmt.Println(a, c)
	const (
		d = iota
		e
		f
	)
	//专业写法,表示给a赋值0 b在a的基础上+1,c在b的基础上+1
	fmt.Println(d, e, f)
	const (
		g    = iota       // 0
		h    = iota       // 1
		i, j = iota, iota //2 2 按换行递增
	)
	fmt.Println(g, h, i, j)

}
