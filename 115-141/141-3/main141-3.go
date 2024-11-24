// package main

// import "fmt"

//	func main() {
//		for i := 11; i <= 100; i++ {
//			if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
//				fmt.Println(i)
//			}
//		}
//	}
package main

import (
	"fmt"
)

// isPrime 判断一个数是否为素数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// printPrimesAndSum 输出100以内的所有素数，每行显示5个，并计算总和
func printPrimesAndSum() {
	var count, sum int
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			count++
			sum += i
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\n100以内所有素数的和为: %d\n", sum)
}

func main() {
	printPrimesAndSum()
}
