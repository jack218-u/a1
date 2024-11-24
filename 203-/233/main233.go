/*
思路分析
把记账软件的功能,封装到一个结构体中,然后调用该结构体的方法,来实现记账,显示明细。结构体的名字FamilyAccount
在通过main方法中,创建一个结构体FamilyAccount实例,实现记账即可
代码不需要重写,只需要重新组织即可
*/
package main

import (
	"abc/203-/233/utils"
	"fmt"
)

func main() {
	fmt.Println("这是一个面向对象的方式完成~~")
	utils.NewfamilyAccount().MainMenu()
}
