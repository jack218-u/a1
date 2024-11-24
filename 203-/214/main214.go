//接口基本介绍
/*
interface类型可以定义一组方法,但是这些不需要实现。并且接口不能包含任何变量，到某个自定义类型(比如结构体Phone)要使用的时候，
再根据具体情况把这些方法写出来(实现--需要接口中所有的方法实现才算实现)
*/
//基本语法
/*
package main
type 接口类型 interface{
//声明了两个没有实现的方法
	方法1(参数列表) (返回值列表)  //比如这里写Start() 后面不能加fmt.Println()类似的语句
	方法2(参数列表) (返回值列表)
	...
}
实现接口所有方法
type 接口类型1 struct{}
//如果接口类型中有多个方法，那么实现就需要多个函数把接口中的方法都实现
 func (变量名1 接口1) 方法1(参数列表) (返回值列表){
	//方法体
 }
 func (变量名1 接口1) 方法2(参数列表) (返回值列表){
	//方法体
 }
 ...
 type 接口类型n struct{}
 func (变量名2 接口类型n) 自定义方法(变量名 接口类型) (返回值列表){
	方法1
	方法2
 }
 自定义方法n为接收方法,将其他实现了接口类型方法的结构体放入到此方法中即可


 func main(){
	变量名3 :=接口类型n{}
	变量名4 = 接口类型1{}
	变量名3.自定义方法{变量名4}
 }
其中变量名1包含接口类型中所有方法，这个变量实现了这个接口
*/
/*
小结说明
1.接口中所有方法都没有方法体,即接口的方法都是没有实现的方法.接口体现了程序设计的多态和高内聚低耦合的思想
2.Golang中的接口,不需要显示的实现(即不需要指定变量实现).只要一个变量,含有接口类型中的所有方法,
那么这个变量就实现了这个接口.因此Golang中没有implement关键字
*/
/*
接口的应用场景
比如:要制造飞机,只需要把飞机的功能/规格定下来即可,然后让别人实现
type interface interface {
	Len() int
	Less(i, j int) bool
}
func Sort (data Interface)
只要实现Interface接口中的Len() int 和 Less(i, j int) bool方法即可,Sort函数就可以排序了
*/
//接口注意事项和细节

// 1.接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
package main

import "fmt"

type AInterface interface {
	say()
}

type B struct {
}

func (b B) say() {
	fmt.Println("golang")
}

type integer int

func (i integer) say() {
	fmt.Println("integer say i =", i)
}

type BInterface interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello")
}
func (m Monster) say() {
	fmt.Println("Monster say")
}

type DInterface interface {
	test01()
}
type EInterface interface {
	test02()
}
type CInterface interface {
	DInterface
	EInterface
	test03()
}

//如果需要实现C接口,需要实现DInterface和EInterface接口中的方法以及C接口中自己的方法test03()
type Router struct {
}

func (r Router) test01() {
	fmt.Println("Router test01")
}
func (r Router) test02() {
	fmt.Println("Router test02")
}
func (r Router) test03() {
	fmt.Println("Router test03")
}

type F interface {
}
type GInterface interface {
	hello01()
	hello02()
}
type HInterface interface {
	hello01()
	hello03()
}
type JInterface interface {
	GInterface
	HInterface
}

func main() {
	//2. 接口中所有的方法都没有方法体,即是没有实现的方法
	//3. 在Golang中,一个自定义类型需要将某个接口的所有方法都是实现,我们说这个自定义类型实现了该接口
	//4. 一个自定义类型只有是实现了某个接口,才能将该自定义类型的实例(变量)赋给接口类型.如代码74行
	//5. 只要是自定义数据类型,就可以实现接口,不仅仅是结构体类型.如代码74 76 87行
	//6. 一个自定义类型可以是实现多个接口,63,80,86,89,102,103,104,105行,但是注意的是,
	//a2变量实现了AInterface接口,可以执行a2.say,不能执行a2.Hello;同理b2变量也一样
	//7. golang接口中不能有任何变量
	//8. 一个接口(比如C接口)可以继承多个别的接口(比如D、E接口),如果变量要实现C接口,需要实现D、E接口中的方法和自己本身的方法
	var a AInterface = B{} //a指向了B实例,B实例实现了say() 实现了AInterface接口,B结构体内容可以随便写
	a.say()
	var i AInterface = integer(10)
	i.say()
	var a2 AInterface = Monster{}
	var b2 BInterface = Monster{}
	a2.say()
	b2.Hello()
	var a3 CInterface = Router{}
	//CInterface中有D和E接口的继承，Router实现了DInterface和EInterface接口方法test01()test02(),
	//以及CInterface中的test03(),所以CInterface的变量a3,实现了CInterface接口
	a3.test03()
	//9.  interface类型默认是一个指针(引用类型),如果没有对interface初始化就使用那么会输出nil
	//参考main213.go中50和51行phone和camera的引用
	//10. 空接口interface{}没有任何方法,意味这可以接收任意一种数据类型,所以所有类型都实现了空接口,
	//即，我们可以把任何一个变量赋值给空接口.比如代码150行
	var stu Router
	var c CInterface = stu
	c.test01()
	var f F = stu
	fmt.Println(f)
	var f2 interface{} = stu //也可以直接写空接口类型 interface{}
	var num1 float64 = 8.8
	var str1 string = "listening raining"
	f2 = num1 //因为f2 f都是空接口类型的变量，可以将任意类型的数据赋值给它们
	f = str1
	fmt.Println(f2, f)
	//11. 一个接口中可以继承多个其他接口，但是继承的多个接口中不能出现同名的方法名,如代码128中出现的2个hello01()方法
	//会提示重复定义的错误

}
