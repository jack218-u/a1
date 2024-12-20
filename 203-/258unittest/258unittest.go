package main

import "fmt"

/*
传统测试方法缺点:
1. 不方便,需要再main函数中去调用,这样就需要去修改main函数,如果现在项目正在运行,就可能去停止项目
2. 不利于管理,因为当我们测试多个函数或多个模块时，都需要写在main函数,不利于我们管理和清晰我们思路

Go语言中自带一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试,testing框架和
其他语言中的测试框架类似,可以基于这个框架写针对相应函数的测试用例,也可以给予该框架写相应的压力测试
用例.通过单元测试,可以解决如下问题:
1. 确保每个函数是可运行,并切运行结果是正确的
2. 确保写出来的代码性能是好的
3. 单元测试能及时的发现程序设计或实现的逻辑错误,使问题及早暴露,便于问题的定位解决,而性能测试的重点在于
发现程序设计上的一些问题,让程序在高并发的情况下还能保持稳定
*/
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}
func main() {
	//传统测试方法,就是在main函数中使用看看结果是否正确
	res := addUpper(10) //1.+10 = 50
	if res != 55 {
		fmt.Println("测试不通过")
	} else {
		fmt.Println("测试通过")
	}
}
