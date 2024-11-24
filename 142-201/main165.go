package main

import "fmt"

func main() {
	//golang中的两种查找
	// 1. 顺序查找
	// 2. 二分查找（该数组是有序的）
	// 有一个数列：三国演义 水浒传 西游记 红楼梦  数列2: "曹操","刘备","孙权"
	// 猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	//思路
	//1. 定义一个数组，三国演义 水浒传 西游记 红楼梦，是字符串数组
	//2. 从控制台接收一个名字，依次比较，如果发现有，提示
	// 方法1.
	novels := [...]string{"三国演义", "水浒传", "西游记", "红楼梦"}
	var name string
	fmt.Println("输入要查找的小说名")
	fmt.Scanln(&name)
	for i := 0; i < len(novels); i++ {
		if name == novels[i] {
			fmt.Printf("找到%v,下标为%v\n", name, i)
			break
		} else if i == len(novels)-1 { //这里必须要加上if i == len(novels)-1，表示最后一次也没找到，否则后面会循环多次没找到的提示
			// 此处表示输出循环查找的次数
			fmt.Printf("没有找到%v\n", name)
		}
	}
	// 方法2.
	heros := [...]string{"曹操", "刘备", "孙权"}
	var name2 string
	fmt.Println("输入三国人物名称")
	fmt.Scanln(&name2)
	index := -1
	for i := 0; i < len(heros); i++ {
		if name2 == heros[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标为%v", name2, index)
	} else {
		fmt.Printf("没有找到%v", name2)
	}
}
