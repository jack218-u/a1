package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	catChan := make(chan *Cat, 10)
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "jerry", Age: 19}
	catChan <- &cat1
	catChan <- &cat2
	//取出
	cat11 := <-catChan
	cat22 := <-catChan
	fmt.Printf("catChan的长度=%v catChan的容量=%v\n", len(catChan), cap(catChan))
	fmt.Println(cat11, cat22)
}
