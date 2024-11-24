package main

import "fmt"

func main() {
	//map是key-value数据结构，又称为字段或者关联数组，类似其他编程语言的集合。
	//map基本语法 var map变量名 map[keytype]valuetype
	//1. map中的keytype可以是很多类型，比如bool，数字，string，指针，channel，还可以是只包含前面几个类型的接口、结构体、数组
	//通常为int、string
	//注意：slice,map还有function不可以，因为这几个没法用==来判断，keytype主要是用来做判断
	//2. map中的valuetype的类型和key基本一样，通常为数字（整数，浮点数），string,map,struct,如果是map就是多重map
	//map的举例声明：
	//var name map[string]string
	//var name map[string]int
	//var name map[int]string
	//var name map[string]map[string]string
	//name := map[string]string
	//注意:map声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
	//map声明的注意事项
	var a map[string]string
	//1. 使用map前需要先make，就是给map分配空间，否则不能赋值使用
	//make使用语法 变量名 = make(需要make的类型,size IntergerType),其中size可以省略，则分配一个小的其实大小
	//2. key不能重复，否则覆盖，value可以重复，只要key不一样
	a = make(map[string]string, 10)
	a["No.1"] = "宋江"
	a["No.2"] = "吴用"
	a["No.1"] = "武松" //覆盖宋江
	a["No.3"] = "吴用"
	a["No.4"] = "九纹龙史进"
	a["No.5"] = "神机军师朱武"
	a["No.6"] = "跳涧虎陈达"
	a["No.7"] = "白花蛇杨春"
	fmt.Println(a)
	//map的使用方式
	//1. 先声明，然后make
	//声明，这是map=nil，未被初始化
	var cities map[string]string
	//make(map[string]string,10)
	cities = make(map[string]string, 10)
	fmt.Println(cities)
	//2. 声明的时候直接make
	var cities2 = make(map[string]string)
	fmt.Println(cities2)
	//3. 声明的同时直接赋值
	var cities3 map[string]string = map[string]string{"No.4": "成都", "No.1": "北京"}
	fmt.Println(cities3)
	a1 := map[string]string{"张三": "18"} //简洁版，直接用类型推导方式，key和value之间用英文:隔开
	a1["李四"] = "19"                     //赋值语法： 变量名[key值]=value值
	fmt.Println(a1)
	//演示一个key-value的案例，存放3个学生信息，每个学生有name和sex信息
	//思路：map[string]map[string]
	//map中的value也是map的情况，需要对此value做make
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 3) //这里的make不能少
	stuMap["stu01"]["name"] = "tom"
	stuMap["stu01"]["sex"] = "男"
	stuMap["stu01"]["addr"] = "北京长安街"
	stuMap["stu02"] = make(map[string]string, 3) //这里的make不能少
	stuMap["stu02"]["name"] = "mary"
	stuMap["stu02"]["sex"] = "女"
	stuMap["stu02"]["addr"] = "上海黄浦江"
	fmt.Println(stuMap)                  //输出学生表信息
	fmt.Println(stuMap["stu01"])         //输出其中一个学生信息
	fmt.Println(stuMap["stu01"]["addr"]) //输出其中一个学生的地址信息
}
