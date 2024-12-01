package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 string `json:"name"`
	Num2 string
}

func (s Cal) GetSub(n1, n2 int) {
	fmt.Printf("tom完成了结构体运算,%v-%v=%v\n", n1, n2, n1-n2)
}
func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	typ := reflect.TypeOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %dfields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field%d的值为%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field%d的tag值为%v\n", i, tagVal)
		}
	}
	val.Method(0).Call([]reflect.Value{reflect.ValueOf(8), reflect.ValueOf(3)})
	/*
		以上一行代码也可以写成:
			var params []reflect.Value
			params = append(params, reflect.ValueOf(8))
			params = append(params, reflect.ValueOf(3))
			val.Method(0).Call(params)
	*/
}
func main() {
	var testCal = Cal{
		Num1: "a",
		Num2: "b",
	}
	TestStruct(testCal)

}
