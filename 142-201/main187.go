package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func main() {
	var p1 Person2
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person2 = &p1 //p2是指向p1的指针类型

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name) //p2.Name其实应该是(*p2).Name，底层已经自动帮忙加上了，与下面代码一样的含义
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name)
	fmt.Printf("p1的地址%p\n", &p1)
	fmt.Printf("p2的地址%p,p2的值%p\n", &p2, p2) //p2的值等于*p1，因为p2是指向p1的指针类型,p2的地址则是它自己本身的指针
	// fmt.Println(*p2.Age) //不能这样子写，会报错，因为.的优先级比*要高，除非Age字段的类型为指针
}
