package main

import (
	"abc/203-/customerManager/model"
	"abc/203-/customerManager/service"
	"fmt"
)

type customerView struct {
	//定义必要字段
	key  string //接收用户输入
	loop bool   //表示是否循环显示菜单
	//增加一个字段service
	service *service.Service //表示service包下的service结构体类型
}

// 显示所有的客户信息
func (p *customerView) list() {
	//首先获取到当前所有客户信息(在切片中)
	customers := p.service.List()
	//显示
	fmt.Println("----------------------------客户列表----------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------------------客户列表完成--------------------------")
	fmt.Println()
}

// 得到用户的输入,信息构建新的客户,并完成添加
func (p *customerView) add() {
	fmt.Println("----------------------------添加客户----------------------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//ID号没有让用户输入,id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email) //重点！！！
	//调用
	if p.service.Add(customer) {
		fmt.Println("----------------------------添加成功----------------------------")
		fmt.Println()
	} else {
		fmt.Println("----------------------------添加失败----------------------------")
		fmt.Println()
	}
}

// 得到用户的输入id,删除该id对应的客户信息
func (cv *customerView) delete() {
	fmt.Println("----------------------------删除客户----------------------------")
	fmt.Println("请选择待删除客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃操作
	}
	fmt.Println("请确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		//调用service包中Service的delete方法
		if cv.service.DeleteById(id) {
			fmt.Println("----------------------------删除成功----------------------------")
		} else {
			fmt.Println("----------------------删除失败,该ID号不存在----------------------")
		}
	}
}

// 修改用户信息
func (p *customerView) update() {
	fmt.Println("----------------------------修改客户----------------------------")
	id := -1
	fmt.Println("请输入待修改用户的ID:")
	fmt.Scanln(&id)
	index := p.service.FindById(id)
	if index == -1 {
		fmt.Println("----------------------修改失败,该ID号不存在----------------------")
		return
	}
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	if p.service.Update(id, name, gender, age, phone, email) {
		fmt.Println("----------------------------修改成功----------------------------")
	} else {
		fmt.Println("----------------------------修改失败----------------------------")
	}
}

// 退出软件
func (cv *customerView) exit() {
	fmt.Println("确认退出? Y/N")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		} else {
			fmt.Println("你的输入有误,请重新输入 Y/N")
		}
	}
	if choice == "Y" || choice == "y" {
		cv.loop = false //这里修改了loop,MainMenu菜单中的loop值也跟着变了，因为是同一个指针结构体
	}
}

// 显示主菜单
func (cv *customerView) MainMenu() {
	for {
		fmt.Println("------------------------客户信息管理软件------------------------")
		fmt.Println("                          1 添加客户")
		fmt.Println("                          2 修改客户")
		fmt.Println("                          3 删除客户")
		fmt.Println("                          4 客户列表")
		fmt.Println("                          5 退出")
		fmt.Println("请选择(1-5):")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("您的输入有误，请重新输入...")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理软件")
}

func main() {
	//在main函数中，创建一个customerView,并运行显示主菜单
	a := customerView{
		key:  "",
		loop: true,
	}
	//这里完成对customerView结构体的service字段的初始化
	a.service = service.NewService() //此处= 后面的service是包名
	//显示主菜单
	a.MainMenu()
}
