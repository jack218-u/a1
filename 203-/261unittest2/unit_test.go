package main

import "testing"

//测试实例,测试Store方法
func TestStore(t *testing.T) {
	//先创建一个monster1实例
	monster := &Monster1{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store1() //调用Store方法,注意调用的语法,前面需要有和方法绑定的结构体变量
	if !res {
		t.Fatalf("moster.Store1()错误,希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store1()测试成功!")
}
func TestRestore(t *testing.T) {
	//先创建一个Monster1实例,不需要指定字段的值,因为创建的Monster1实例对返回结果无影响
	//返回的res是调用的已经存在的序列化后的字符串,此时Monster1只是一个空结构体用来调用
	//反序列化方法中的bool结果
	var monster = &Monster1{}
	res := monster.Restore1()
	if !res {
		t.Fatalf("monster.Restore1()错误,希望为=%v 实际为=%v", true, res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Restore1()错误,希望为=%v 实际为=%v", "红孩儿", res)
	}
	t.Logf("moster.Restore1()测试成功!")
}
