/*
面向对象编程-多态
1. 基本介绍
变量(实例)具有多种形态。面向对象的第三大特征是多态。在Go语言中，多态特征是通过接口实现的。
可以按照统一的接口来调用不同的实现。这是接口变量就呈现不同的形态。
2. 多态的概念： 多态是面向对象编程中的一个重要概念，它允许一个接口或基类引用多个派生类对象。在 Go 语言中，
多态是通过接口实现的。接口定义了一组方法签名，任何实现了这些方法的类型都可以被视为实现了该接口。
3. 快速入门
在前面的Usb接口案例中，Usb usb，既可以接收手机变量，又可以接收相机变量，就体现了Usb接口的多态特性

接口体现多态特征
1. 多态参数
再前面的Usb接口案例中，Usb usb，既可以接收手机变量，又可以接收相机变量，就体现了Usb接口的多态特性。
2. 多态的返回值
在Go语言中，函数的返回值也可以是接口类型，接口变量接收函数的返回值，就体现了多态性。
3. 多态数组
演示一个案例，给Usb数组中，存放Phone结构体和Camera结构体变量,Phone还有一个特有的方法call(),请遍历Usb数组,
如果是Phone变量，除了调用Usb接口声明的方法外，还需要调用Phone特有方法call()。
*/

package main

import "fmt"

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	name string
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {
}

//编写一个方法working 方法，接收一个Usb接口变量
//只要实现了Usb接口(所谓实现usb接口，就是指实现了usb接口声明的所有方法)
func (c Computer) Working(usb Usb) {
	//变量usb会更具传入的实参，来判断到底是PHone还是Camera，属于多态&动态绑定的理论,usb叫做多态参数
	usb.Start()
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里体现多态数组，一个接口多态数组中，既可以存放Phone，又可以存放Camera
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
}
