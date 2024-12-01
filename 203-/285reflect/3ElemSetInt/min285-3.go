package main

import (
	"fmt"
	"reflect"
)

// 通过反射,修改 num int的值
// 修改student值
func reflectTest01(b interface{}) {
	//获取reflect.Value值
	rVal := reflect.ValueOf(b)
	//看看rVal的Kind是
	fmt.Printf("rVal的kind=%v\n", rVal.Kind()) //传入的是指针数据,输出的类型就是ptr
	rVal.Elem().SetInt(20)                    //先用Elem()方法获取指针指向的值,然后修改
	/*
		Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
		如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值
	*/
}
func main() {
	var num int = 10
	reflectTest01(&num)
	fmt.Println("num=", num) //reflectTest01(&num)修改num的值
}
