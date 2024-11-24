package main

import (
	"fmt"
)

// 使用结构体struct
// 1. 定义一个cat结构体，将cat的各个字段/属性信息，放入到cat结构体进行管理
type Cat struct { //结构体名称首字母大写，才能被其他包引用，小写的只能在本包使用
	Name   string //各个字段名首字母大写，该字段才能被其他包引用，小写的属于私有，只能在本包使用
	Age    int
	Color  string
	Scores [3]int
	res    map[string]string
}

func main() {
	//2 .创建一个cat变量,数据类型为之前创建的结构体，Cat为之前创建的结构体
	//2.1 并对其赋值,语法为：结构体变量名.结构体字段名 =
	var cat1 Cat
	fmt.Println(cat1) //默认为空，string为空字符串，int为0，因为没有赋值
	fmt.Printf("cat1的数据类型%T,cat1的地址=%p\n", cat1, &cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	//3. 输出变量属性
	fmt.Println(cat1)
	fmt.Println("猫的信息如下: ")
	fmt.Println("name", cat1.Name)
	fmt.Println("age", cat1.Age)
	fmt.Println("color", cat1.Color)
	fmt.Printf("cat1的数据类型%T,cat1的地址=%p\n", cat1, &cat1)
	//结构体和结构体变量(实例)的区别和联系
	//1. 结构体是自定义的数据类型，代表一类事物，比如上面例子中的Cat
	//2. 结构体变量(实例)是具体的，实际的，代表一个具体的变量，比如上面例子中的cat1
	//3. 结构体和结构体变量在内存中值类型，不是引用类型
	//如何声明结构体
	// type 结构体名称 struct {
	//    field1 type
	//    field2 type
	//  }
	//1. 结构体字段=属性=field
	//2. 字段是结构体的一个组成部分，一般是基本数据类型、数组，也可以是引用类型。比如我们前面定义毛结构体的Name string就是属性
	//字段&属性的注意事项：
	//1. 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前面讲的一样：
	//1.1. 布尔类型是false，数值型是0，字符串是""
	//1.2. 数组类型的默认值和它的元素类型相关，比如score [3]int 则为[0,0,0]
	//1.3. 指针,slice,和map的零值都是nil，即还没有分配空间,需要make以后才可以使用
	var p1 Person //Person结构体在main函数下面
	fmt.Println(p1)
	//没有赋值之前，以下三种类型都能匹配nil
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok1")
	}
	if p1.map1 == nil {
		fmt.Println("ok1")
	}
	//结构体字段为slice的赋值需要make以后才能使用,同样使用map也需要make
	p1.slice = make([]int, 3)
	p1.slice[0] = 100
	fmt.Println(p1)
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println(p1)
	//2. 不同的结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个
	//结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500
	monster2 := monster1
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
	monster2.Name = "青牛精" //monster2修改后不影响monster1，结构体默认为值拷贝
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)

}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64        //数组
	ptr    *int              //指针
	slice  []int             //切片    使用需要make
	map1   map[string]string //map切片 使用需要make
}
type Monster struct {
	Name string
	Age  int
}
