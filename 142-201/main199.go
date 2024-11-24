// 面向对象编程应用实例步骤
// 1. 声明(定义)结构体
// 2. 编写结构体的字段
// 3. 编写结构体的方法
package main

import "fmt"

//学生案例：
//1. 编写一个student结构体，包含name,gender,age,id,score字段，类型分别为string,string,int,int,float64类型
//2. 结构体中声明一个say方法，返回string类型，方法返回信息中包含name,gender,age,id,score字段的信息
//3. 在main主函数中创建一个student实例，调用say方法，输出学生信息
type Student2 struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student2) say() string {
	infoStr := fmt.Sprintf("name=%v,gender=%v,age=%v,id=%v,score=%v",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}
func main() {
	//测试，创建一个Student实例变量
	var stu = Student2{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     10001,
		score:  90.5,
	}
	fmt.Println(stu.say())
}
