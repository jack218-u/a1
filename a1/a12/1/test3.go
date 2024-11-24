package main

import "fmt"

func defertest(i int) (r int) {
	defer func() {
		r += i
	}()

	return 2

}
func main() {
	fmt.Println(defertest(1))
}
