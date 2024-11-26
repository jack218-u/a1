package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (p *Monster) Store() string {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("序列化失败,err=%v", err)
	}
	return string(data)
}

func (p *Monster) Restore(jsondata string) *Monster {
	var c Monster
	err := json.Unmarshal([]byte(jsondata), &c)
	if err != nil {
		fmt.Printf("反序列化失败,err=%v", err)
	}
	return &c
}
func main() {
	a := &Monster{"张三", 18, "太极拳"}
	b := a.Store()
	fmt.Println(b)
	d := &Monster{}
	e := d.Restore(b)
	fmt.Println(*e)
}
