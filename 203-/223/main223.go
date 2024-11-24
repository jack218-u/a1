package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	a = point
	fmt.Println(a)
	// b = a 不可以
	// b = a.(Point)可以
	b := a.(Point) //类型断言，表示判断a是否指向Point类型的变量，
	//如果是就转成Point类型并赋给b变量，否则就报错
	fmt.Println(b)
	//类型断言的其它案例
	var x interface{}
	var c float32 = 1.1
	x = c //空接口，可以接收任意类型
	// x => float32 [使用类型断言]
	y := x.(float32)
	/*
		当一个变量被声明为 interface{} 类型时，它实际上可以存储任何类型的值。
		虽然 x 当前存储的是 float32 类型的值，但在运行时，x 可能会存储其他类型的值。
		因此，为了确保从 interface{} 类型中提取的值确实是预期的类型，需要使用类型断言。
	*/
	fmt.Printf("y的类型是%T 值是=%v\n", y, y)
	//在类型断言的时候，如果类型不匹配，就会报panic,因此进行类型断言时候，要确保原来空接口指向的，就是要断言的类型
	//在断言时,可以带上检测机制，如果成功就ok,否则也不要报panic
	d, ok := x.(float64)
	if ok { //简易方法ok是布尔类型
		fmt.Println("转换成功")
		fmt.Println(d)
	} else {
		fmt.Println("类型转换失败")
	}

}
