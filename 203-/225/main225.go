package main

import "fmt"

//定义Student类型结构体
type Student struct {
	name string
	age  int
}

//编写一个函数，判断输出的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) { //x.(type)获取x的数据类型
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型,值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型,值是%v\n", index, x) //类型断言不仅能判断已有的数据类型,也可以判断自定义的类型
		case *Student:
			fmt.Printf("第%v个参数是指针类型,值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型未知,值是%v\n", index, x)
		}
	}
}
func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	stu1 := Student{"tom", 18}
	stu2 := &Student{"jack", 19}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)

	// stu1 := Student{}
	// stu2 := &Student{}
	// fmt.Printf("%T,%T", stu1, stu2)
}
