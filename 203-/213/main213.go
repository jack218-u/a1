// 接口快速入门
package main

import "fmt"

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
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
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//关键点：Phone和Camera实现了Usb接口，所以可以作为参数传递给Working()
	computer.Working(phone)
	computer.Working(camera)
}
