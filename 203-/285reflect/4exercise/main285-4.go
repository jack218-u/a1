package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	rVal := reflect.ValueOf(v)
	rTyp := reflect.TypeOf(rVal)
	rKin := rVal.Kind()
	fmt.Println(rVal, rTyp, rKin)
	iV := rVal.Interface()
	num := iV.(float64)
	fmt.Println(num)
	var str string = "tom"
	// fs := reflect.ValueOf(str)
	fs1 := reflect.ValueOf(&str)
	fs1.Elem().SetString("jack")
	fmt.Println(str)
}
