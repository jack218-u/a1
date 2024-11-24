package main

import "fmt"

func shixingoldtower(m int) {
	//打印实心金字塔
	for i := 1; i <= m; i++ {
		//j表示每层打印几个*
		//j = i 表示每层打印*的个数和层数一样
		for k := 1; k <= m-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
		}
		fmt.Println() //每次j的内层循环完成后到输入i的时候换行
	}
}
func kongxingoldtower(m int) {
	//打印空心金字塔
	for i := 1; i <= m; i++ {
		//j表示每层打印几个*
		//j = i 表示每层打印*的个数和层数一样
		for k := 1; k <= m-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == m { //i==m 即最后一行，三者任一一个成立都输入*
				fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() //每次j的内层循环完成后到输入i的时候换行
	}
}
func shixinlingxing(m int) {
	//打印实心菱形
	for i := 1; i <= m; i++ {
		//j表示每层打印几个*
		//j = i 表示每层打印*的个数和层数一样
		if i <= (m+1)/2 {
			for k := 1; k <= (m+1)/2-i; k++ {
				fmt.Print(" ")
			}
			for j := 1; j <= 2*i-1; j++ {
				fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
			}
		} else {
			for k := 1; k <= i-(m+1)/2; k++ {
				fmt.Print(" ")
			}
			for j := 1; j <= 2*(m+1-i)-1; j++ {
				fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
			}
		}
		fmt.Println() //每次j的内层循环完成后到输入i的时候换行
	}
}
func kongxinlingxing(m int) {
	for i := 1; i <= m; i++ {
		//j表示每层打印几个*
		//j = i 表示每层打印*的个数和层数一样
		if i <= (m+1)/2 {
			for k := 1; k <= (m+1)/2-i; k++ {
				fmt.Print(" ")
			}
			for j := 1; j <= 2*i-1; j++ {
				if j == 1 || j == 2*i-1 {
					fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
				} else {
					fmt.Print(" ")
				}
			}
		} else {
			for k := 1; k <= i-(m+1)/2; k++ {
				fmt.Print(" ")
			}
			for j := 1; j <= 2*(m+1-i)-1; j++ {
				if j == 1 || j == 2*(m+1-i)-1 || i == m {
					fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println() //每次j的内层循环完成后到输入i的时候换行
	}
}
func MultiplyTable(Num int) {
	for i := 1; i <= Num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	var m int
	var n int
	var o int
	var p int
	var q int
	fmt.Println("输入实心金字塔的层数")
	fmt.Scanln(&m)
	shixingoldtower(m)
	fmt.Println("输入空心金字塔的层数")
	fmt.Scanln(&n)
	kongxingoldtower(n)
	fmt.Println("输入实心菱形塔的层数-请输入奇数")
	fmt.Scanln(&o)
	shixinlingxing(o)
	fmt.Println("输入空心菱形塔的层数-请输入奇数")
	fmt.Scanln(&p)
	kongxinlingxing(p)
	fmt.Println("输入九以内的乘法表-数字低于九")
	fmt.Scanln(&q)
	MultiplyTable(q)
}
