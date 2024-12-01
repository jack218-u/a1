package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
func TestStruct(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白象精")
	for i := 0; i < num; i++ {
		fmt.Printf("Field%d:值为=%v\n", i, val.Elem().Field(i))
	}
	fmt.Printf("struct has %d fields\n", num)
	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	val.Elem().Method(0).Call(nil)
}
func main() {
	var a Monster = Monster{
		Name:  "黄狮子",
		Age:   408,
		Score: 92.8,
	}
	//说明Marshal就是通过反射获取到struct的tag值
	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))
	TestStruct(&a)
	fmt.Println(a) //传入的是黄狮子,但是是指针类型,TestStruct函数中修改成了白象精,然后原始的a也跟着变了
}
