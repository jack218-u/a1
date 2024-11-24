// 面向对象编程应用实例步骤
// 1. 声明(定义)结构体
// 2. 编写结构体的字段
// 3. 编写结构体的方法
package main

import "fmt"

//盒子案例：
//1. 创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽、高，长宽高要始终从终端获取
//2. 声明一个方法获取立方体的体积
//3. 在main函数中，创建一个Box结构体变量，并调用方法，打印给定尺寸的立方体的体积
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolume() float64 {
	return box.len * box.width * box.height
}
func main() {
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volume := box.getVolume()
	fmt.Printf("体积=%.2f", volume)
}
