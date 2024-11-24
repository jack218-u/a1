package main

import "fmt"

func modify(map1 map[int]int) {
	map1[1] = 900
}

//定义一个学生结构体,语法type 结构体名称 struct {}
type stu struct {
	name string
	age  int
	addr string
}

func main() {
	//map使用细节
	//1. map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
	map1 := make(map[int]int, 2) //容量为2，但是容量满后会自动扩容
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1) //结果map1[10]=900,说明map是引用类型
	//2. map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长"键值对"
	//3. map的value也经常使用struct类型，更适合管理复杂的数据(比前面value用一个map更好管理)，比如value为student结构体
	//map的key为学生的学号，是唯一的；map的值为结构体，包含学生的名字，年龄，地址
	student := make(map[string]stu)
	//创建两个学生
	stu1 := stu{name: "tom", age: 18, addr: "北京"}
	stu3 := stu{"mary", 20, "南京"} //也可以直接输入结构体内的值
	student["No.1"] = stu1
	student["No.3"] = stu3
	fmt.Println(student)
	//遍历学生信息
	for k, v := range student {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的姓名是%v\n", v.name)
		fmt.Printf("学生的年龄是%v\n", v.age)
		fmt.Printf("学生的地址是%v\n", v.addr)
		fmt.Println()
	}
}
