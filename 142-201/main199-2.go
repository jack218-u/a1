// 面向对象编程应用实例步骤
// 1. 声明(定义)结构体
// 2. 编写结构体的字段
// 3. 编写结构体的方法
package main

import "fmt"

//景区门票案例：
//1. 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20，其他情况门票免费
//2. 请编写Visitor结构体，根据年龄端决定能够购买的门票价格并输出
type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) showPrice() {
	if v.Age >= 90 || v.Age <= 8 {
		fmt.Println("考虑到安全问题,就不要玩了")
		return //结束整个程序，不考虑下面的if，或者也可以不用return,后面的if之前加上else
	}
	if v.Age >= 18 {
		fmt.Printf("游客的姓名为%v,年龄为%v,收费20元\n", v.Name, v.Age)
	} else {
		fmt.Printf("游客的姓名为%v,年龄为%v,免费\n", v.Name, v.Age)
	}
}

func main() {
	var v Visitor
	for {
		fmt.Println("请输入你的姓名")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
