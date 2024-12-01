package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string
	Age   int
	Score float32
}

func (m *Monster) GetSum(a, b int) int {
	return a + b
}

func main() {
	var m Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
	}

	// 获取结构体实例的 reflect.Value
	val := reflect.ValueOf(&m)

	// 获取方法
	method := val.MethodByName("GetSum")

	// 准备参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) // 第一个参数
	params = append(params, reflect.ValueOf(40)) // 第二个参数

	// 调用方法
	res := method.Call(params)

	// 获取返回值
	sum := res[0].Int() // 返回值是 int 类型，所以用 Int() 方法获取
	fmt.Println("sum =", sum)
}
