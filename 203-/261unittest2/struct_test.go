package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster1 struct {
	Name  string
	Age   int
	Skill string
}

func (p *Monster1) Store1() bool {
	//先序列化
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("marshal err =%v", err)
		return false
	}
	//保存文件
	filepath := "D:/practice/unittest/monster.ser"
	err = os.WriteFile(filepath, data, 0666) //如果该路径下没有文件,会自动生成该文件
	if err != nil {
		fmt.Printf("write file err =%v", err)
		return false
	}
	return true
}

func (p *Monster1) Restore1() bool {
	//1.先从文件中,读取序列化的字符串
	filepath := "D:/practice/unittest/monster.ser" //定义一个序列化后的文件
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadFile err =%v", err)
		return false
	}
	//2.使用读取到的data []byte,对其反序列化
	err = json.Unmarshal(data, p)
	if err != nil {
		fmt.Printf("Unmarshal err =%v", err)
		return false
	}
	return true
}
