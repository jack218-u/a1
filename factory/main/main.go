package main

import (
	"abc/factory/model"
	"fmt"
)

func main() {
	/*
		var stu = model.Student{ //已经引入了外部的model包,里面有结构体Student
			Name:  "tom",
			Score: 78.9,
		}
		fmt.Println(stu)
	*/
	stu := model.NewStudent("tom", 88.9) //指向结构体的指针
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
