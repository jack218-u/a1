package main

import (
	"testing" //引入testing包
)

// 编写要给测试用例
func TestGetSub(t *testing.T) { //Test开头,后面的第一个字母要大写,类似TestXxx,是一个内置测试函数
	//调用getSub()函数来自同一个文件夹下面的cal.go中,
	//因为两个代码的package包名是相同的,所以可以直接调用,不用import
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub(10,3)执行错误,期望值=%v 实际值=%v\n", 7, res)
	}
	//如果正确,输出日志
	t.Logf("getSub(10,3)执行正确")
	//测试验证命令在当前路径下执行 go test -v,而非go run命令,
}
func TestGetMul(t *testing.T) { //Test开头,后面的第一个字母要大写,类似TestXxx,是一个内置测试函数
	//调用getSub()函数来自同一个文件夹下面的cal.go中,
	//因为两个代码的package包名是相同的,所以可以直接调用,不用import
	res := getMul(2, 6)
	if res != 12 {
		t.Fatalf("getSub(10,3)执行错误,期望值=%v 实际值=%v\n", 12, res)
	}
	//如果正确,输出日志
	t.Logf("getSub(2,6)执行正确")
	//测试验证命令在当前路径下执行 go test -v,而非go run命令,
}

/*
在 Go 语言中，go test 命令会自动发现并运行当前目录及其子目录中
所有以 _test.go 结尾的文件中的测试函数。

然而，如果测试代码没有被执行或没有输出预期的日志，可能有以下几种原因：
1.包名不一致：确保 subtest.go 和 cal.go 文件的包名完全一致。如果包名
不一致，即使文件在同一目录下，也无法直接调用函数。
2.测试函数命名不规范：测试函数必须以 Test 开头，并且第一个参数必须是
 *testing.T 类型。确保 TestGetSub 函数的命名和参数符合要求。
3.函数未定义：确保 getSub 函数在 cal.go 文件中定义，并且函数签名正确。
4.测试文件路径问题：确保 subtest.go 文件在 go test 命令执行的目录中，
或者在子目录中被正确识别。
5.测试输出被抑制：默认情况下，go test 命令只会显示失败的测试结果。
使用 -v 标志可以显示详细的测试输出，包括成功和失败的测试。
*/
/*
总结:
1.测试用例文件名必须以_test.go结尾。比如cal_test.go
2.测试用例函数必须以Test开头,一般来说就是Test+被测试的函数名(自定义,首字母要大写),比如TestGetSub,不能是
TestgetSub,然后{}中的执行方法引出测试的具体函数(外部包中的函数)
3.TestAddUpper(t *testing.T)的形式参数必须是*testing.T,具体含义参考官方文档
4.一个测试用例文件中,可以有多个测试用例函数,比如TestGetSub、TestAddUpper(可以放在一个或多个go文件中)
5.运行测试用例指令
a)cmd>go test[如果运行正确,无日志,错误时,会输出日志]
b)cmd>go test -v[运行正确或错误,都会输出日志]
6.当出现错误时,可以使用t.Fatalf来格式化输出错误信息,并退出程序
7.t.Logf可以输出相应日志
8.测试用例函数,并没有放在main函数中,也执行了,因为是被测试框架自动调用了,这是测试用例的方便之处
9.PASS表示测试用例运行成功,FAIL表示测试用例运行失败
10.测试单个方法,一般要带上被测试的源文件,除非两者在一个文件中,
语法:go test -v 测试方法文件 被测试的源文件,语法中两个文件的顺序任意
a)测试方法和被测试内容分别放在两个文件,比如go test -v cal_test.go cal.go
b)测试方法和被测试内容在一个文件：比如go test -v sub2_test.go
11.如果只想测试其中一个函数用例,语法：go test -v -run=TestGetMul,=后面写要测试的函数名
*/
