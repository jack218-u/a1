package main

import "fmt"

// 定义Usb接口
type Usb interface {
	Start()
	Stop()
}

// 定义Phone结构体并实现Usb接口
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Printf("%s 手机开始工作...\n", p.Name)
}

func (p Phone) Stop() {
	fmt.Printf("%s 手机停止工作...\n", p.Name)
}

// 定义Camera结构体并实现Usb接口
type Camera struct {
	Model string
}

func (c Camera) Start() {
	fmt.Printf("%s 相机开始工作...\n", c.Model)
}

func (c Camera) Stop() {
	fmt.Printf("%s 相机停止工作...\n", c.Model)
}

// 计算机结构体
type Computer struct {
}

// 编写一个方法working，接收一个Usb接口变量
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 定义一个函数，返回Usb接口类型
func GetUsbDevice(deviceType string) Usb {
	if deviceType == "phone" {
		return Phone{Name: "iPhone"}
	} else if deviceType == "camera" {
		return Camera{Model: "Canon EOS R5"}
	}
	return nil
}

func main() {
	// 创建计算机实例
	computer := Computer{}

	// 获取Usb设备
	phone := GetUsbDevice("phone")
	camera := GetUsbDevice("camera")

	// 通过接口变量接收返回值，体现多态性
	computer.Working(phone)
	computer.Working(camera)
}
