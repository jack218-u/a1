// 对于结构体的序列化,如果我们希望序列化后的key的名字是我们想要的,可以给struct字段添加tag标签
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	name string // 字段用小写,json是外置包,在没有封装的情况下,无法引用,Marshal函数只能舍弃
	Age  int
}
type Monster2 struct {
	Name string
	Age  int
}
type Monster3 struct {
	Name string `json:"name"` //希望序列化后的Name是小写输出,此处打上tag
	Age  int
}

// 演示将结构体,map,切片进行序列化
// 对结构体序列化
func test1() {
	//演示
	monster := Monster{
		name: "牛魔王",
		Age:  500,
	}
	data, err := json.Marshal(&monster) //在没有封装的情况下,无法引用结构体中的首字母小写的字段,Marshal函数只能舍弃
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("妖怪 :%v------", string(data)) //data是[]byte类型,查看需要转成string类型
	fmt.Println("说明:无封装时,无法引用结构体中的首字母小写的字段,Marshal函数只能舍弃结构体中的name字段")
}
func test2() {
	//演示
	monster := Monster2{
		Name: "牛魔王",
		Age:  500,
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("妖怪 :%v------", string(data)) //data是[]byte类型,查看需要转成string类型
	fmt.Println("说明:正常的大写的Name输出")
}
func test3() {
	//演示
	monster := Monster3{
		Name: "牛魔王",
		Age:  500,
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("妖怪 :%v------", string(data)) //data是[]byte类型,查看需要转成string类型
	fmt.Println("说明:结构体字段打上tag标签,输出序列化后的key名字是tag标签的值,Name变成name")
}

func main() {
	test1()
	test2()
	test3()
}
