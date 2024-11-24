package main

import "fmt"

func fbn(n uint8) uint8 {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
func main() {
	fmt.Println(fbn(6))
}
