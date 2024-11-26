// json反序列化是指,将json字符串反序列化成对应的数据类型(比如结构体、map、切片)的操作
package main

import (
	"encoding/json"
	"fmt"
)

func testMap() string {
	//定义一个map 使用map前需要make
	a := make(map[string]interface{})
	a["name"] = "红孩儿~~~"
	a["age"] = 30
	a["address"] = "火云洞"
	//将a这个map进行序列化
	data, err := json.Marshal(a) //map本身是指针类型,这里不需要再传指针类型
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	// fmt.Printf("a map序列化后 =%v\n", string(data)) //序列化后的结果排序和赋值的顺序无关
	return string(data)
}

// 演示将json字符串反序列化成结构体
// 定义一个结构体
type Monster struct {
	Name    string
	Age     int
	Address string
}

func unmarshalStruct() {
	//str表示序列化后的json字符串，在项目开发中,是通过网络传输或者读取文件获取到的,以下只是为了演示
	// str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}" //双引号里面再有双引号为了避免当成是字符串的结束
	str := testMap() //通过程序获取的json字符串,就可以不用加引号,中间也不需要做转义处理
	var monster Monster
	//对序列化后的字符串str进行反序列化,然后传给&monster,只有通过&monter才能改变monster的值
	//反序列化后赋值的数据类型,需要和需要反序列化的数据在被序列化之前的数据类型一致,即monster变量的数据类型
	//要和str中反序列化前的数据类型一致,比如两者都是结构体,name结构体中的字段类型需要一致,数量也需要一致
	//特殊情况下,比如序列化前是map类型,其序列化后的数据,可以反序列化成结构体类型,其字段一样就可以
	err := json.Unmarshal([]byte(str), &monster) //反序列化成bytes类型切片,然后传递给&monster
	if err != nil {
		fmt.Printf("反序列化失败 err= %v\n", err)
	}
	fmt.Printf("反序列化成功,变成结构体变量,monster=%v\n", monster)
	fmt.Printf("反序列化成功,结构体变量也可以单独取出结构体其中一个字段,monster.name=%v\n", monster.Name)
	// fmt.Printf("反序列化成功,结构体变量也可以单独取出结构体其中一个字段,monster.address=%v\n", monster.address)
}

// 演示将json字符串,反序列化成map,反序列化时候map不需要make,因为make已经被封装到了Unmarshal函数中
func unmarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}           //定义的map类型应该和字符串序列化之前的类型一样
	err := json.Unmarshal([]byte(str), &a) //Unmarshal需要修改a的值,所以需要用指针类型
	if err != nil {
		fmt.Printf("反序列化失败 err= %v\n", err)
	}
	fmt.Printf("反序列化成功,变成map变量,a=%v\n", a)
	fmt.Printf("反序列化成功,变成map变量,也可以单独取出map的其中一个key对应的value,a[\"name\"]=%v\n", a["name"])
}

// 演示将json字符串,反序列化成slice,反序列化时候slice不需要make,因为make已经被封装到了Unmarshal函数中
func unmarshalSlice() {
	//多行字符串用``; 如str
	//如果用"",需要换行的时候,多行之间用+拼接,每一行收尾要有",中间的"前面加\转义,如str2
	str := `[{"address":"北京","age":7,"name":"jack"},
	{"address":["墨西哥","夏威夷"],"age":12,"name":"tom"}]`
	str2 := "[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":12,\"name\":\"tom\"}]"
	var slice []map[string]interface{} //定义的slice类型应该和字符串序列化之前的类型一样
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化失败 err= %v\n", err)
	}
	fmt.Printf("反序列化成功,变成slice变量,slice=%v\n", slice)
	fmt.Printf("反序列化成功,slice变量也可以单独取出其中一个元素,slice[0]=%v\n", slice[0])
	fmt.Printf("反序列化成功,slice变量也可以单独取出其中一个元素的字段,slice[0][\"address\"]=%v\n", slice[0]["address"])

	var slice2 []map[string]interface{}
	err2 := json.Unmarshal([]byte(str2), &slice2)
	if err2 != nil {
		fmt.Printf("反序列化失败 err= %v\n", err2)
	}
	fmt.Printf("反序列化成功,变成slice变量,slice2=%v\n", slice2)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
