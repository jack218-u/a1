// 多重继承
package main

import "fmt"

type A struct {
	Name string
	Age  int
}
type B struct {
	Name string
	Age  int
}
type C struct {
	A
	B
}
type Stu struct {
	A
	int
	n int // 一个结构体有int类型的匿名字段，就不能有第二个，如果需要有多个int字段，就需要给int字段指定名字
}

//多重继承说明
//如果一个struct中嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，
//从而实现多重继承，如上面的type C类型
//为了代码的可读性，最好不要在struct中嵌套多个匿名结构体
func main() {
	stu := Stu{}
	stu.Name = "tom"
	stu.Age = 10
	stu.int = 80
	stu.n = 90
	fmt.Println(stu)
	var c C
	fmt.Println(c)

}
