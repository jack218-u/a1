// 此案例为223案例的缩减版，不适用封装(工厂模式)，FamilyAccount结构体和内部的字段的首字母都是大写，所以外部可以访问到
package utils

import "fmt"

type FamilyAccount struct { //结构体首字母大写，可以被外部包引用，为共有
	//声明一个字段，保存接收用户输入的选项
	Key string //字段首字母大写，可以被外部包引用，为共有
	//声明一个字段,控制是否退出for
	Loop bool
	//定义账户的余额[]
	Balance float64
	//每次收支的金额
	Money float64
	//每次收支的说明
	Note string
	//定义一个字段，记录是否有收支的行为
	Flag bool
	//收支详情使用字符串来记录
	//当有收支时候，只需要对Details进行拼接处理即可
	Details string
}

//编写一个工厂模式的构造方法,返回一个FamilyAccount实例
// func NewFamilyAccount() *FamilyAccount {
// 	return &FamilyAccount{
// 		Key:     "",
// 		Loop:    true,
// 		Balance: 10000.0,
// 		Money:   0.0,
// 		Note:    "",
// 		Flag:    false,
// 		Details: "收支\t收支金额\t账户余额\t说明\n",
// 	}
// }
//因为FamilyAccount结构体，首字母大写，字段首字母也大写，所以不需要使用工厂模式(中间引用函数)

//将显示明细写成一个方法
func (p *FamilyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-------------------")
	if p.Flag { //Flag定义的是bool类型的数据,默认是true，此处省略了== true
		fmt.Println(p.Details)
	} else { //else提示的是false
		fmt.Println("当前没有收支明细...来一笔吧")
	}
}

//将登记收入写成一个方法,和*FamilyAccount绑定
func (p *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&p.Money)
	p.Balance += p.Money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&p.Note)
	//将这个情况，拼接到Details变量
	//收入     11000       1000      有人发红包
	p.Details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", p.Money, p.Balance, p.Note)
	p.Flag = true
}

//将登记支出写成一个方法,和*FamilyAccount绑定
func (p *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&p.Money)
	//需要做一个判断
	if p.Money > p.Balance {
		fmt.Println("余额不足")
		return
	}
	p.Balance -= p.Money //修改账户余额
	fmt.Println("本次支出说明:")
	fmt.Scanln(&p.Note)
	p.Details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", p.Money, p.Balance, p.Note)
	p.Flag = true
}

//将退出写成一个方法,和*FamilyAccount绑定
func (p *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误,请重新输入 y/n")
	}
	if choice == "y" {
		p.Loop = false
	}
}

// 给结构体绑定相应的方法
// 显示主菜单
func (p *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&p.Key)
		switch p.Key {
		case "1":
			p.showDetails()
		case "2":
			p.income()
		case "3":
			p.pay()
		case "4":
			p.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !p.Loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}
