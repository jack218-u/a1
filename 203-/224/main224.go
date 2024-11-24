// 在前面的Usb接口案例做改进:给Phone结构体增加一个特有的方法call(),当Usb接口接收的是Phone变量时，还需要调用call方法
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

func (p Phone) Call() { //Phone结构体中添加了Call()方法,独有的，其他结构体没有
	fmt.Println("手机打电话...")
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
	//如果usb是指向Phone结构体变量，则还需要调用call方法
	//类型断言...
	if phone, ok := usb.(Phone); ok { //phone在这里是变量可以换成任何字符，ok是bool类型,简写为ok等于true
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里体现多态数组，一个接口多态数组中，既可以存放Phone，又可以存放Camera
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"尼康"}
	//遍历usbArr
	//Phone还有一个特有的方法call(),请遍历Usb数组，如果是phone变量，
	//除了调用Usb接口声明的方法外，还需要调用phone特有方法call.=》类型断言
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	//以上for range循环方法遍历数组，执行Computer的Working方法
	//computer结构体working方法绑定，working中的参数类型是Usb接口，
	// fmt.Println(usbArr)
	//以上调用注意事项：
	/*
		1.接口类型：Usb 是一个接口类型，它声明了Start和Stop两个方法。
		任何实现了这两个方法的结构体都可以被视为 Usb 类型。

		2.多态性：在 usbArr 中，虽然实际存储的是 Phone和Camera 的实例，
		但它们都实现了 Usb 接口。因此，这些实例可以被赋值给Usb类型的变量

		3.类型统一：在 usbArr 数组中，尽管实际存储的是不同类型的对象（如 Phone 和 Camera），
		但这些对象都实现了同一个接口（Usb 接口）。因此，它们可以被统一地视为 Usb 类型的变量。
		这种统一的类型处理方式使得我们可以使用相同的接口方法来操作这些不同类型的对象，从而实现多态性

		4.类型断言：在 computer.Working(v) 中，v是一个Usb类型的变量，但是它的具体类型是 Phone 或 Camera。
		为了调用Phone或Camera 的特有方法，我们需要使用类型断言来将 v 转换为 Phone 或 Camera 类型。
	*/

}
