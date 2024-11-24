package main

import "fmt"

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (p *Monkey) climbing() {
	fmt.Println(p.Name, "生来会爬树")
}

// littlemonkey结构体
type LittleMonkey struct {
	Monkey //继承了Monkey的所有属性
}

func (p *LittleMonkey) Flying() {
	fmt.Println(p.Name, "通过学习会飞")
}
func (p *LittleMonkey) Swimming() {
	fmt.Println(p.Name, "通过学习会游泳")
}

//让littlemonkey实现flying和swimming方法，在设计上没有破坏原有的struct结构体，同时通过接口进行了扩展

/* 实现接口vs继承
1. 接口和继承解决的问题不同
a.继承的价值主要在于：解决代码的复用性和可维护性
b.接口的价值主要在于：设计，设计好各种规范(方法)，让其他自定义类型去实现这些方法。
2. 接口比继承更加灵活
接口比继承更加灵活，继承是满足is -a 的关系,而接口值需要满足like -a 的关系
3. 接口在一定成都上实现代码解耦
*/
func main() {
	//创建一个littlemonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
	//总结：
	//1. 当a结构体继承了b结构体，a结构体就拥有了b结构体的所有字段和方法，并且可以直接使用
	//2. 当a结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可,因此我们可以认为，实现接口是对继承机制的补充
}
