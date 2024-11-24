package main

import "fmt"

func main() {
	//map切片，切片的数据类型如果是map,则我们称为slice of map，map切片，这样使用则map个数就可以动态变化了
	//案例：使用一个map来记录monster信息name和age，即一个monster对应一个map，并且妖怪的个数可以动态的增加=>map切片
	//1. 声明一个map切片
	monsters := make([]map[string]string, 2) //声明一个map切片，长度为2，准备放入两个妖怪
	//2. 增加一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2) //map的key值为name和age
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2) //map的key值为name和age
		monsters[1]["name"] = "蜈蚣精"
		monsters[1]["age"] = "400"
	}
	//下面写法越界
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2) //map的key值为name和age
	// 	monsters[2]["name"] = "蝎子精"
	// 	monsters[2]["age"] = "400"
	// }
	// map切片扩容，用切片的append函数，可以动态的增加monster
	//1. 先定义一个新的monster信息
	newMonster := map[string]string{
		"name": "新的妖怪-青牛精",
		"age":  "200",
	}
	//2. 将新的monster加到原来map中
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

}
