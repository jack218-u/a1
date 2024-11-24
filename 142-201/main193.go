// 方法的调用和传参机制原理
package main

import "fmt"

//说明,方法的调用和传参机制和函数基本一样,不一样的地方是方法调用时,会将调用方法的变量,当作实参传递给方法
//1. 在通过一个变量去调用方法时,其调用机制和函数一样
//2. 不一样的地方是,变量调用方法时,该变量本身也会作为一个参数传递到方法(如果变量是值类型,则进行值拷贝;如果变量是引用类型,则进行地址拷贝)
//举例:编写一个程序,要求如下:
//a. 声明一个结构体Circle,字段为radius
//b. 声明一个方法area和Circle绑定,可以返回面积
//c. 画出area执行过程+说明
type Circle struct {
	radius float64 //radius为圆的半径
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	var c1 Circle
	c1.radius = 4.0
	res := c1.area()
	fmt.Println("圆的面积=", res)
}
