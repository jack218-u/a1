package main

import (
	"fmt"
	"strings"
)

// 闭包函数
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func makeSuffix2(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func makeSuffix3(name string, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}
func main() {
	f := makeSuffix(".png")
	fmt.Println("文件名处理后=", f("winter.png"))
	f1 := makeSuffix2(".doc")
	fmt.Println("文件名处理后=", f1("summer"))
	fmt.Println("文件名处理后=", makeSuffix3("ppt", "August"))
}
