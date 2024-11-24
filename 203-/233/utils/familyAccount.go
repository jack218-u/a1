package utils

import "fmt"

type familyAccount struct {
	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段,控制是否退出for
	loop bool
	//定义账户的余额[]
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义一个字段，记录是否有收支的行为
	flag bool
	//收支详情使用字符串来记录
	//当有收支时候，只需要对details进行拼接处理即可
	details string
}

//编写一个工厂模式的构造方法,返回一个familyAccount实例--对familyAccount结构体做封装,family为小写不能被访问，需要用到函数调用
func NewfamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t收支金额\t账户余额\t说明\n",
	}
}

//将显示明细写成一个方法
func (p *familyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-------------------")
	if p.flag { //flag定义的是bool类型的数据,默认是true，此处省略了== true
		fmt.Println(p.details)
	} else { //else提示的是false
		fmt.Println("当前没有收支明细...来一笔吧")
	}
}

//将登记收入写成一个方法,和*familyAccount绑定
func (p *familyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&p.money)
	p.balance += p.money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&p.note)
	//将这个情况，拼接到details变量
	//收入     11000       1000      有人发红包
	p.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", p.money, p.balance, p.note)
	p.flag = true
}

//将登记支出写成一个方法,和*familyAccount绑定
func (p *familyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&p.money)
	//需要做一个判断
	if p.money > p.balance {
		fmt.Println("余额不足")
		return
	}
	p.balance -= p.money //修改账户余额
	fmt.Println("本次支出说明:")
	fmt.Scanln(&p.note)
	p.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", p.money, p.balance, p.note)
	p.flag = true
}

//将退出写成一个方法,和*familyAccount绑定
func (p *familyAccount) exit() {
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
		p.loop = false
	}
}

// 给结构体绑定相应的方法
// 显示主菜单
func (p *familyAccount) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&p.key)
		switch p.key {
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
		if !p.loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}
