package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共用了多少次", count)
}
