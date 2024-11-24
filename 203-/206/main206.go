package main

import (
	"abc/203-/206/model"
	"fmt"
)

func main() {
	//创建account变量
	// account := model.NewAccount("jzh11111", "999999", 40) //用小括号不是中括号
	// if account != nil {
	// 	fmt.Println("创建成功=", *account)
	// } else {
	// 	fmt.Println("创建失败")
	// }
	// account.Deposite(100, "999999") //存款测试
	// account.WithDraw(50, "999999") //取款测试
	// account.Query("999999")        //查询测试
	//以上为测试代码
	var str, pwd string
	var balance float64
	fmt.Println("请输入需要创建的账户名")
	fmt.Scanln(&str)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	fmt.Println("请输入首次存款的金额")
	fmt.Scanln(&balance)
	account := model.NewAccount(str, pwd, balance) //测试账号创建是否成功
	if account != nil {
		fmt.Println("创建成功=", *account)
	} else {
		fmt.Println("创建失败")
	}
	for {
		var op int
		fmt.Println("请输入操作:1.存款 2.取款 3.查询 4.退出")
		fmt.Scanln(&op)
		switch op {
		case 1:
			var m float64
			var p string
			fmt.Println("请输入需要存款的金额和密码,以空格隔开")
			fmt.Scanln(&m, &p)
			account.Deposite(m, p)
		case 2:
			var m float64
			var p string
			fmt.Println("请输入需要取款的金额和密码,以空格隔开")
			fmt.Scanln(&m, &p)
			account.WithDraw(m, p)
		case 3:
			var p string
			fmt.Println("请输入密码")
			fmt.Scanln(&p)
			account.Query(p)
		case 4:
			return
		default:
			fmt.Println("输入有误")
		}
	}
}
