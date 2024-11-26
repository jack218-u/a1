package main

import (
	"testing" //引入testing包
)

// 编写要给测试用例,去测试addUpper是否正确
func TestAddUpper(t *testing.T) { //Test开头,后面的第一个字母要大写,类似TestXxx,是一个内置测试函数
	//调用addUpper()函数来自同一个文件夹下面的cal.go中,
	//因为两个代码的package包名是相同的,所以可以直接调用,不用import
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper(10)执行错误,期望值=%v 实际值=%v\n", 55, res)
	}
	//如果正确,输出日志
	t.Logf("addUpper(10)执行正确")
	//测试验证命令在当前路径下执行 go test -v,而非go run命令
}
