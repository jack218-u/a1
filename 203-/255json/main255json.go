/*
[{"name":"tom","age":18,"address":["上海","杭州"],"hobby":["足球","篮球"]},
{"name":"jack","age":19,"address":["北京","成都"]}]
*/

/*
JSON 数据由键值对组成，键和值之间用冒号 : 分隔，每对键值对之间用逗号 , 分隔。
键必须是字符串，值可以是多种类型，包括字符串、数字、布尔值、数组、对象（嵌套的 JSON）、null。
*/

/*
json.Marshal(&monster) 将 monster 结构体序列化为 JSON 格式的 []byte。
string(data) 将 []byte 转换为字符串，以便在控制台中打印出可读的 JSON 字符串。
*/

package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string `json:"monster_name"` //反射机制,将Name改成希望的格式
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 演示将结构体,map,切片进行序列化
// 对结构体序列化
func testStruct() {
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2019-09-09",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	data, err := json.Marshal(&monster) //为了好操作传指针类型
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster =%v\n", string(data)) //data是[]byte类型,查看需要转成string类型
}

// interface{} 类型是一种空接口类型，它可以表示任何类型的值。
// 这意味着你可以将任何类型的值赋给 interface{} 类型的变量
// 将map进行序列化
func testMap() {
	//定义一个map 使用map前需要make
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"
	//将a这个map进行序列化
	data, err := json.Marshal(a) //map本身是指针类型,这里不需要再传指针类型
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map序列化后 =%v\n", string(data)) //序列化后的结果排序和赋值的顺序无关
}

// 对切片序列化,切片有很多map,属于切片类型的map,[]map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}
	//使用map前需要先make
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 12
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)
	//将切片进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err==%v\n", err)
	}
	fmt.Printf("slice序列化后的结果=%s\n", data) //也可以用string(data), %v表示
}

// 对基本数据类型序列化,对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err==%v\n", err)
	}
	fmt.Printf("num1序列化后的结果=%s\n", data)
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
