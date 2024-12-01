package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct { // 检查值是否为结构体
		fmt.Println("expect struct")
		return
	}
	num := val.NumField() // 获取结构体的字段数量
	/*
		func (v Value) NumField() int
		返回v持有的结构体类型值的字段数，如果v的Kind不是Struct会panic
	*/

	fmt.Printf("struct has %d fields\n", num) // 输出结构体的字段数量,Monster中定义几个就输出几个
	for i := 0; i < num; i++ {
		fmt.Printf("Field%d:值为=%v\n", i, val.Field(i))
		//获取到struct标签,注意需要通过reflect.Type
		tagVal := typ.Field(i).Tag.Get("json") // 获取字段的JSON标签,Tag是字段的标签
		if tagVal != "" {                      // 如果标签不为空
			fmt.Printf("Field%d:tag为=%v\n", i, tagVal)
		}
		/*
		   		func (v Value) Field(i int) Value
		   返回结构体的第i个字段（的Value封装）。如果v的Kind不是Struct或i出界会panic
		*/
		/*
			func (tag StructTag) Get(key string) string
			Get方法返回标签字符串中键key对应的值。如果标签中没有该键，会返回""。
			如果标签不符合标准格式，Get的返回值是不确定的。
		*/
	}
	numOfMethod := val.NumMethod() // 获取结构体的方法数量
	/*
	   	func (v Value) NumMethod() int
	   返回v持有值的方法集的方法数目
	*/
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 遍历方法并打印方法名
	for i := 0; i < numOfMethod; i++ {
		// method := val.Method(i)
		methodName := typ.Method(i).Name
		fmt.Printf("Method%d: 名称为=%s\n", i, methodName)
	}

	//var params[]reflect.Value
	val.Method(1).Call(nil) // 调用结构体的第二个方法，不传递参数
	//调用结构体的第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) // 添加第一个参数
	params = append(params, reflect.ValueOf(40)) // 添加第二个参数
	res := val.Method(0).Call(params)            // 调用方法并传递参数
	fmt.Println("res=", res[0].Int())            //返回结果,返回的结果是
}
func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
		//此处没有定义Sex字段和值,就是默认为空,输出空格
	}
	TestStruct(a)
}
