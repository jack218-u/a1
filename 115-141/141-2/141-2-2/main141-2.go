//以下是阿里AI工具写的代码

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNumberGame() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成1到11之间的随机数
	target := rand.Intn(11) + 1
	var guess int
	var attempts int

	fmt.Println("猜数字游戏开始！数字在1到11之间，你有10次机会。")

	for attempts = 0; attempts < 10; attempts++ {
		fmt.Printf("请输入你的猜测（第%d次尝试）: ", attempts+1)
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("输入错误，请输入一个有效的数字。")
			continue
		}

		if guess == target {
			switch attempts {
			case 0:
				fmt.Println("你真是个天才")
			case 1, 2:
				fmt.Println("你很聪明")
			case 3, 4, 5, 6, 7, 8:
				fmt.Println("一般般")
			case 9:
				fmt.Println("可算猜对了")
			}
			return
		} else if guess < target {
			fmt.Println("太小了！")
		} else {
			fmt.Println("太大了！")
		}
	}

	fmt.Println("说你点啥好呢")
}

func main() {
	guessNumberGame()
}
