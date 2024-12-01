/*
思路:
1.创建管道用来存放1--2000;启用一个协程来执行写入操作
2.启动8个协程,执行读取操作,将1+2+3+...+n结果写入另一个管道中
3.将结果写入到映射中
*/
package main

import (
	"fmt"
	"sync"
)

// 计算1到n的和
func sum(n int) int {
	return n * (n + 1) / 2
}

func main() {
	// 创建两个 channel
	numChan := make(chan int)
	resChan := make(chan int)

	// 启动一个 goroutine 将 1 到 2000 的数字放入 numChan 中
	go func() {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	// 使用 WaitGroup 等待所有 worker 协程完成
	var wg sync.WaitGroup

	// 启动 8 个 worker 协程
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range numChan {
				resChan <- sum(n)
			}
		}()
	}

	// 等待所有 worker 协程完成
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// 遍历 resChan，显示结果
	results := make([]int, 0, 2000)
	for res := range resChan {
		results = append(results, res)
	}

	// 打印结果
	for i, res := range results {
		fmt.Printf("res[%d] = %d\n", i+1, res)
	}
}
