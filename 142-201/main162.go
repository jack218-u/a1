package main

import "fmt"

func BubbleSort(arr *[5]int) {
	//对数列元素做从小到大排序
	fmt.Println("排序前arr=", (*arr))
	//以下仅是为了演示效果的手动循环代码，详细可以参考BubbleSort2的代码
	//完成第一轮排序(外层第一次)
	// tmp := 0 //临时变量用于交换，也可以用a,b = b,a更简单的方式
	for j := 0; j < 4; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// //交换
			// tmp = (*arr)[j]
			// (*arr)[j] = (*arr)[j+1]
			// (*arr)[j+1] = tmp
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
		}
	}
	fmt.Println("第一次排序后arr=", (*arr))
	//完成第二轮排序(外层第二次)
	// tmp := 0 //临时变量用于交换，也可以用a,b = b,a更简单的方式
	for j := 0; j < 3; j++ { //第二轮排序 j < 3
		if (*arr)[j] > (*arr)[j+1] {
			// //交换
			// tmp = (*arr)[j]
			// (*arr)[j] = (*arr)[j+1]
			// (*arr)[j+1] = tmp
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
		}
	}
	fmt.Println("第二次排序后arr=", (*arr))
	//完成第三轮排序(外层第3次)
	// tmp := 0 //临时变量用于交换，也可以用a,b = b,a更简单的方式
	for j := 0; j < 2; j++ { //第三轮排序 j < 2
		if (*arr)[j] > (*arr)[j+1] {
			// //交换
			// tmp = (*arr)[j]
			// (*arr)[j] = (*arr)[j+1]
			// (*arr)[j+1] = tmp
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
		}
	}
	fmt.Println("第三次排序后arr=", (*arr))
	//完成第四轮排序(外层第4次)
	// tmp := 0 //临时变量用于交换，也可以用a,b = b,a更简单的方式
	for j := 0; j < 1; j++ { //第四轮排序 j < 1
		if (*arr)[j] > (*arr)[j+1] {
			// //交换
			// tmp = (*arr)[j]
			// (*arr)[j] = (*arr)[j+1]
			// (*arr)[j+1] = tmp
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
		}
	}
	fmt.Println("第四次排序后arr=", (*arr))
}
func BubbleSort2(arr *[5]int) {
	//对数列元素做从小到大排序
	fmt.Println("排序前arr2=", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] { //如果是从大到小排序，此处的>改成<
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println("排序前arr2=", (*arr))
}
func BubbleSort3(arr *[5]int) {
	//对数列元素做从大到小排序
	fmt.Println("排序前arr3=", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] < (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println("排序前arr3=", (*arr))
}
func main() {
	//排序是将一组数据，按照指定的顺序进行排列的过程
	//排序的分类:
	//1. 内部排序
	//指将需要处理的所有数据都加载到内部存储器中进行排序
	//包括（交换式排序法、选择式排序法和插入式排序法）；
	//2. 外部排序法：
	//数据量过大，无法全部加载到内存中，需要借助外部存储进行排序
	//包括（合并排序法和直接合并排序法）
	//交换式排序属于内部排序法，是运用数据值比较后，按照判断规则对数据位置进行交换，以达到排序的目的
	//交换式排序分为两种：
	//1. 冒泡排序法（Bubble Sort)
	//2. 快速排序法（Quick Sort）
	//冒泡排序法一共经过len(arr)-1轮比较，每一轮比较次数再逐渐减少，当发现下标靠前的数比下标靠后的数打的时候，进行交换
	//冒泡排序法的代码思路1.将最大数放到最后；2.将第二大的数放到倒数第二位；以此类推；然后组成多重循环函数，先写内层循环然后外层
	arr := [5]int{50, 40, 30, 20, 10}
	//将数组传递给一个函数，完成排序
	BubbleSort(&arr)
	arr2 := [5]int{5, 4, 3, 2, 1}
	BubbleSort2(&arr2)
	arr3 := [5]int{6, 7, 8, 9, 10}
	BubbleSort3(&arr3)
	fmt.Println("main中arr3的排序", arr3) //主函数中的排序也改变了
}
