package main

import "fmt"

func main() {
	//定义二维数组，用于保存三个班，每个班五名同学成绩
	//并求出每个班级平均分，以及所有班级平均分
	//1. 定义一个二维数组
	var scores [3][5]float64
	//2. 循环输入成绩
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d个班级的第%d个同学的成绩", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	// fmt.Println(scores)
	//3. 遍历输出成绩后的二维数组，统计平均分
	totalsum := 0.0 //定义一个变量，统计所有班级的总分
	for i := 0; i < len(scores); i++ {
		sum := 0.0 //定义一个变量，用于累计各个班级的总分
		for j := 0; j < (len(scores[i])); j++ {
			sum += scores[i][j]
		}
		totalsum += sum
		fmt.Printf("第%d个班级总分为%v,平均分为%v\n",
			i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v,所有班级的平均分为%v\n",
		totalsum, totalsum/15)
}
