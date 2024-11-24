package main

import "fmt"

func main() {
	//二维数组使用方式
	//1. 先声明再赋值，语法：var 数组名 [a1][a2]类型，a1表示二维数组中共有几个一维数组，a2表示每个一维数组中有几个元素,比如var arr [2][3]int
	//数组赋值中元素赋值方法，比如，数组名[b1][b2],b1表示表示二维数组中排序为b1的一维数组（排序从0开始）
	//中的的排序为b2（排序为从0开始）的元素的值
	var arr2 [2][3]int
	arr2[1][2] = 10
	fmt.Println(arr2)
	fmt.Printf("arr2[0]的地址%p\n", &arr2[0]) //表示所在一维数组的第一个元素的地址和arr2[0][0]地址一致
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1]) //每个一维数组中有3个元素，所以arr2[0]和arr2[1]相差24个字节

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])
	//2. 直接初始化
	// 声明:var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
	// 赋值（有默认值，比如int类型就是0)
	//说明，二维数组在声明的时候有四种写法和一维数组类似
	// 1. var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
	// 2. var 数组名 [大小][大小]类型 = [...][大小]类型{{初值...},{初值...}}
	// 3. var 数组名 = [大小][大小]类型{{初值...},{初值...}}
	// 4. var 数组名 = [...][大小]类型{{初值...},{初值...}} //根据数组后面的值，自动判断数组多少行，每行多少个
	// 5. 数组名 := [...][大小]类型{{初值...},{初值...}}
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)
	var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	arr5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr4=", arr4)
	fmt.Println("arr5=", arr5)

}