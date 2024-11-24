// 1.创建程序，在model包中定义Account结构体：在main函数中体会golang的封装性
// 2. account结构体要求具有字段：账号(长度在6--10之间)、余额(必须>20)、密码(必须是6位)
// 3. 通过setXX的方法给account的字段赋值,通过GetXXX方法获取account的字段值
// 4. 在main函数中测试
package model

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额必须大于20")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

//存款
func (p *account) Deposite(money float64, passwd string) {
	if passwd != p.pwd {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("输入的金额不正确")
		return
	}
	p.balance += money
	fmt.Printf("存款成功,账户%v余额为%.2f\n", p.accountNo, p.balance)
}

//取款
func (p *account) WithDraw(money float64, passwd string) {
	if passwd != p.pwd {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 || money > p.balance {
		fmt.Println("输入的金额不正确")
		return
	}
	p.balance -= money
	fmt.Printf("取款成功,账户%v余额为%.2f\n", p.accountNo, p.balance)
}

//查询余额
func (p *account) Query(passwd string) {
	if passwd != p.pwd {
		fmt.Println("密码不正确")
		return
	}
	fmt.Printf("账户%v余额为%.2f\n", p.accountNo, p.balance)
}
