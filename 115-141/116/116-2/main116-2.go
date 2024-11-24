package main

import "fmt"

func f(n int) int {
	//已知f(1)=3,当n>1时，f(n)=2*f(n-1)+1,求f(5)和f(10)的值
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
func main() {
	fmt.Println("f(5)=", f(2))
	fmt.Println("f(10)=", f(10))
}
