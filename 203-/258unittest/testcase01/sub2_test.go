package main

import (
	"testing" //引入testing包
)

func sub(n1 int, n2 int) int {
	res := n1 + n2
	return res
}

// 编写要给测试用例
func TestGetSubLocal(t *testing.T) { //Test开头,后面的第一个字母要大写,类似TestXxx,是一个内置测试函数
	//调用getSub()函数来自同一个文件夹下面的cal.go中,
	//因为两个代码的package包名是相同的,所以可以直接调用,不用import
	res := sub(3, 6)
	if res != 9 {
		t.Fatalf("sub(3,6)执行错误,期望值=%v 实际值=%v\n", 9, res)
	}
	//如果正确,输出日志
	t.Logf("sub(3,6)执行正确")
	//测试验证命令在当前路径下执行 go test -v,而非go run命令,
}
