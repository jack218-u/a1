/*
项目开发流程
项目需求说明
项目的界面
项目代码实现
*/
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type AccountData struct {
	Balance float64
	Details string
	Flag    bool
}

func loadData() (AccountData, error) {
	file, err := os.Open("account_data.gob")
	if err != nil {
		if os.IsNotExist(err) {
			return AccountData{}, nil // 如果文件不存在，返回空数据
		}
		return AccountData{}, err
	}
	defer file.Close()

	var data AccountData
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return AccountData{}, err
	}
	return data, nil
}

func saveData(data AccountData) error {
	file, err := os.Create("account_data.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	//声明一个变量，保存接收用户输入的选项
	key := ""
	//声明一个变量,控制是否退出for
	loop := true
	//定义账户的余额[]
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义一个变量，记录是否有收支的行为
	flag := false
	//收支详情使用字符串来记录
	//当有收支时候，只需要对details进行拼接处理即可
	details := "收支\t收支金额\t账户余额\t说明\n"
	// 读取上次的数据
	data, err := loadData()
	if err != nil {
		fmt.Println("读取数据失败:", err)
	} else {
		balance = data.Balance
		details = data.Details
		flag = data.Flag
	}
	//显示这个主菜单
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("-----------------当前收支明细记录-------------------")
			if flag { //flag定义的是bool类型的数据,默认是true，此处省略了== true
				fmt.Println(details)
			} else { //else提示的是false
				fmt.Println("当前没有收支明细...来一笔吧")
			}
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			//将这个情况，拼接到details变量
			//收入     11000       1000      有人发红包
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", money, balance, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			//需要做一个判断
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money //修改账户余额
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", money, balance, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	// 保存当前的数据
	dataToSave := AccountData{
		Balance: balance,
		Details: details,
		Flag:    flag,
	}
	err = saveData(dataToSave)
	if err != nil {
		fmt.Println("保存数据失败:", err)
	}
	fmt.Println("你退出了家庭记账软件的使用")
}
