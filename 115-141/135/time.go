package main

import (
	"fmt"
	"time"
)

func main() {
	//看看日期和时间相关函数和方法的使用
	//1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)
	//2. 通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	//3. 格式化日期时间
	// 方式1：使用Printf或者Sprintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v", dateStr) //返回一个变量dateStr，然后输出变量dateStr
	// 方式2：使用time.Format()方法完成
	// 2006-01-02 15:04:05 数字是固定的
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02")) //只取出年月日
	fmt.Println()
	fmt.Printf(now.Format("15:04:05")) //只取出时分秒
	fmt.Println()
	fmt.Printf(now.Format("01")) //只取出月份
	fmt.Println()
	//时间常量 纳秒 微妙 毫秒 秒 之间换算是1000  分钟 小时 换算是60
	//时间常量的应用，结合sleep
	//需求，每隔1秒钟打印一个数字，打印到100时就退出，间隔时间用时间休眠函数表示
	//需求2，每个0.1秒打印一个数字，打印到100时候就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		// time.Sleep(time.Second)         //休眠1秒
		time.Sleep(time.Millisecond * 10) //休眠0.1秒，毫秒乘以100，不能用秒乘以0.1，编译器过不去
		if i == 2 {
			break
		}
	}
	//7. 获取当前unix时间戳 和unixnano时间戳（作用是可以获取随机数字）
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v", now.Unix(), now.UnixNano())

}
