package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明一个hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2. 声明一个hero结构体切片类型
type HeroSlice []Hero

// 3. 实现Interface接口
// 此案例中是隐式实现,HeroSlice 类型实现了 sort.Interface 接口的 Len, Less, 和 Swap 方法，sort.Sort(a) 可以成功执行
func (hs HeroSlice) Len() int {
	//Len方法返回集合中的元素个数
	return len(hs)
}

// Less方法就是决定你使用什么标准进行排序
// 1.按Hero的年龄从小到大排序,关键点
func (hs HeroSlice) Less(i, j int) bool {
	// return hs[i].Age < hs[j].Age
	//修改成按照名字排序
	// Less方法报告索引i的元素是否比索引j的元素小
	return hs[i].Name < hs[j].Name
}
func (hs HeroSlice) Swap(i, j int) {
	// Swap方法交换索引i和j的两个元素
	hs[i], hs[j] = hs[j], hs[i]
}

// 4.声明一个Student结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//要求将Student的切片,按照Score从大到小排序,课后练习

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对intSlice进行排序
	//1. 冒泡排序
	//2. 系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var a HeroSlice
	for i := 0; i < 10; i++ {
		b := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(11)),
			Age:  rand.Intn(100),
		}
		//将hero添加到heros切片中
		// fmt.Println(b)
		a = append(a, b)

	}
	fmt.Println(a)
	//看看排序前的顺序
	// for _, v := range a {
	// 	fmt.Println(v)
	// }
	//由于 HeroSlice 类型实现了 sort.Interface 接口的 Len, Less, 和 Swap 方法，sort.Sort(a) 可以成功执行
	sort.Sort(a) //调用sort包中的Sort函数对a进行排序,a满足Interface接口中的三个方法
	fmt.Println("------排序后-----")
	fmt.Println(a)
	// fmt.Println(a)
	//看看排序后的顺序
	// for _, v := range a {
	// 	fmt.Println(v)
	// }

}
