package main

import (
	"fmt"
)

func test115(n int) {
	//斐波拉契数列
	if n > 2 {
		n--
		test115(n)
	}
	fmt.Println("n=", n)
}
func main() {
	test115(4)
}
