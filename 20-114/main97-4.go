package main
import "fmt"
func main(){
	//使用for循环完成以下案例写一个程序表示，可以接收一个整数，表示层数，打印出金字塔
	//编程思路
	//1.打印出一个矩形
	/* 
		***
		***
		***
	*/
	//2.打印半个金字塔
	/*
		*     1个*
		**    2个*
		***   3个*
	*/
	//9.打印整个金字塔
	/*  假如总层数是m，i表示第几层，每层空格= m - i
	       *       1层1个* 规律 2 * 层数 -1
	      ***      2层3个* 规律 2 * 层数 -1
		 *****     3层5个* 规律 2 * 层数 -1

		   *       1层1个*  空格  3  
		  ***      2层3个*  空格  2
		 *****     3层5个*  空格  1
		*******    4层7个*  空格  0

		   *       1层1个*  空格  3  
		  ***      2层3个*  空格  2
		 *****     3层5个*  空格  1
		*******    4层7个*  空格  0
		 *****     5层5个*  空格  1
		  ***      6层3个*  空格  2
		   *       7层1个*  空格  3
	*/
	//i<=(m+1)/2时，每层空格数=(m+1)/2 -i,每层的*数=2i-1，有空心空格的情况下，在2i-1处输入*
	//i> (m+1)/2时，每层空格数=i-(m+1)/2, 每层的*数=2(m+1-i)-1，有空心空格的情况下，在2(m+1-i)-1处输入*
	//设置空格数为m
	//i表示层数
	var m int 
	fmt.Println("输出菱形的层数，画出菱形形状")
	fmt.Scanln(&m)
	for i :=1; i <=m; i++ {
		//j表示每层打印几个*
		//j = i 表示每层打印*的个数和层数一样
		if i <= (m+1)/2 {
			for k :=1; k <=(m+1)/2 -i; k++ {
				fmt.Print(" ")
			}
			for j :=1; j <=2 * i - 1; j++ {
				if j == 1 || j == 2*i-1 {
				fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
				} else {
					fmt.Print(" ")
				}
			}
		} else {
			for k :=1; k <=i-(m+1)/2 ; k++ {
				fmt.Print(" ")
			}
			for j :=1; j <=2*(m+1-i)-1; j++ {
				if j == 1 || j == 2*(m+1-i)-1 || i ==m {
				fmt.Print("*") //这里用fmt.Print不用fmt.Println,每次输出不用换行
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println() //每次j的内层循环完成后到输入i的时候换行
	}
}


package main
import "fmt"
func main(){
	var m int 
	fmt.Println("输出菱形的层数，画出菱形形状")
	fmt.Scanln(&m)
	for i :=1; i <=m; i++ {

		if i <= (m+1)/2 {
			for k :=1; k <=(m+1)/2 -i; k++ {
				fmt.Print(" ")
			}
			for j :=1; j <=2 * i - 1; j++ {
				if j == 1 || j == 2*i-1 {
				fmt.Print("*") 
				} else {
					fmt.Print(" ")
				}
			}
		} else {
			for k :=1; k <=i-(m+1)/2 ; k++ {
				fmt.Print(" ")
			}
			for j :=1; j <=2*(m+1-i)-1; j++ {
				if j == 1 || j == 2*(m+1-i)-1 || i ==m {
				fmt.Print("*") 
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println() 
	}
}
