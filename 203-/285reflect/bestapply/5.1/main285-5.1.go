package main

import "fmt"

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
	}
	fmt.Println(a) //这里没有设置Sex性别,则输出结果默认为空""作为输出结果
}
