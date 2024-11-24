// 结构体使用注意事项和细节
package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}
type Rect struct { //Rect此处是矩形的意思
	leftUp    Point //字段类型为Point结构体
	rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point //另一种字段表示方法，此时存放的是指向Point的指针
}

func main() {
	//1. 结构体的所有字段在内存中是连续的
	r1 := Rect{Point{1, 2}, Point{3, 4}} //双重结构体赋值
	//r1有四个int，在内存中是连续分布
	//打印地址
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	fmt.Println(r1)
	fmt.Printf("r1.leftUp.x 地址=%v r1.leftUp.y 地址%v r1.rightDown.x 地址=%v r1.rightDown.y 地址=%v \n",
		r1.leftUp.x, r1.leftUp.y, r1.rightDown.x, r1.rightDown.y)

	//r2有两个 *Point类型字段，这两个*Point类型的本身地址也是也是连续的，但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 本身地址=%p r1.rightDown 本身地址=%p \n", &r2.leftUp, &r2.rightDown)
	//r2指向的地址不一定连续,要看系统运行时如何分配的
	fmt.Printf("r2.leftUp 指向地址=%p r1.rightDown 指向地址=%p \n", r2.leftUp, r2.rightDown)
	//2. 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段(名称、个数和类型)
	var a A
	var b B
	// a = b // 将b赋给a，会提示报错，因为结构体类型不一样，一个是A tpye一个是B type
	a = A(b) // 将b转换为A类型,然后才可以赋值给a
	fmt.Println(a, b)
	//3. 结构体进行type重新定义(相当于取别名)，Golang认为是新的数据类型，但是相互之间可以强转
	var t1 Student1
	var t2 Stu
	// t2 = t1 //虽然t2的类型Stu是取别名Student1,但是被认为是新的类型，所以不能相互转换
	t2 = Stu(t1) //但是将t1转换为Stu类型，然后再赋值给t2，是可以的
	fmt.Println(t2)
	//4. struct的每个字段上，可以写一个tag，该tag可以通过反射机制获取，场景的使用场景就是序列化和反序列化
	//4.1 创建一个Monster变量
	monster := Monster3{"牛魔王", 500, "芭蕉扇"}
	//4.2 将monster变量序列化为json格式字符串，保存到文件中
	jsonStr, err := json.Marshal(monster) //注意monster中的字段名首字母要大写，否则获取不到，因为json来自外部包
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println("jsonStr=", jsonStr) //返回的是bytes类型
	fmt.Println("jsonStr=", string(jsonStr))
}

type A struct {
	Num int
}
type B struct {
	Num int
}
type Student1 struct {
	Name string
	Age  int
}
type Stu Student1 // 取别名
type Monster3 struct {
	Name  string `json:"name"` //json返回值的字段首字母是小写需要按照此操作打tag,注意字段值需要加英文双引号
	Age   int    `json:"age"`  //`json:"name"` 就是struct的tag
	Skill string `json:"skill"`
}
