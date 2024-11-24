/*
思路分析
把记账软件的功能,封装到一个结构体中,然后调用该结构体的方法,来实现记账,显示明细。结构体的名字FamilyAccount
在通过main方法中,创建一个结构体FamilyAccount实例,实现记账即可
代码不需要重写,只需要重新组织即可
*/
package main

import (
	"abc/203-/233-1/utils"
	"fmt"
)

func main() {
	fmt.Println("这是一个面向对象的方式完成~~")
	// utils.NewFamilyAccount().MainMenu()
	a := utils.FamilyAccount{
		Key:     "",
		Loop:    true,
		Balance: 10000.0,
		Money:   0.0,
		Note:    "",
		Flag:    false,
		Details: "收支\t收支金额\t账户余额\t说明\n",
	}
	//需要先对结构体进行初始化，然后使用结构体指针调用方法
	a.MainMenu()
}
