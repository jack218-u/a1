// 创建一个allChan,最多可以存放10个任意类型数据变量,演示写入和读取的用法
package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10) //定义空接口,什么数据都可以存到管道中
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "jerry", Age: 19}
	cat3 := Cat{Name: "blackcat", Age: 20}
	allChan <- cat3
	allChan <- cat2
	allChan <- cat1
	allChan <- 10
	allChan <- "Jack"
	//取出
	<-allChan
	<-allChan
	//希望获得管道中的第三个元素,则先将前两个元素推出
	newCat := <-allChan
	fmt.Printf("netCat=%T, netCat=%v\n", newCat, newCat)
	/*
		cat11 是从allChan中取出的数据，类型为interface{}。
		虽然实际存储的是Cat结构体，但在编译时，Go语言无法确定
		cat11 的具体类型，因此不能直接访问其Name字段,需要使用类型断言
	*/

	/*
		cat, ok := newCat.(Cat)
		if ok {
			fmt.Println("类型断言成功,输出结构体字段", cat.Name)
		} else {
			fmt.Println("类型断言失败,newCat is not a Cat")
		}
	*/
	//以上断言如果能100%确定,也可以简写成
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
	v1 := <-allChan
	v2 := <-allChan
	fmt.Println(newCat, v1, v2)
}
