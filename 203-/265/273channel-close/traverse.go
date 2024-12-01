/*
channel的关闭
使用内置函数close可以关闭channel,当channel关闭后,就不能再向channel写数据了,但是仍然可以从该channel读取数据
channel的遍历
channel支持for--range的方式进行遍历,请注意两个细节
1.在遍历时,如果channel没有关闭,则出现deadlock的错误
2.在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历
*/
package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close
	//这时不能再写入数到channel
	// intChan <- 300
	fmt.Println("ok")
	//当管道关闭后,读取数据是可以的
	n1 := <-intChan
	n2 := <-intChan
	n3 := <-intChan
	n4 := <-intChan
	fmt.Println(n1, n2, n3, n4)
	fmt.Println(len(intChan))
	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}
	//遍历管道不能使用普通的for循环,因为管道的长度会变化
	// for i :=0,i <len(intChan2);i++ {
	// }
	/*
		在遍历时,如果channel没有关闭,则会出现deadlock的错误
		在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历
	*/
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
	fmt.Printf("%T", intChan2)
}
