package model

import "fmt"

type person struct {
	Name string
	age  int //其它包不能直接访问..
	sal  float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age // 将传入的 age 参数值赋给 person 结构体中的 age 字段
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员给一个默认值
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal // 将传入的 sal 参数值赋给 person 结构体中的 sal 字段
	} else {
		fmt.Println("薪水范围不正确..")

	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
