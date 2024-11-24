package main

import (
	"fmt"
	"time"
)

// isFishingOrDrying 判断给定日期是“打鱼”还是“晒网”
func isFishingOrDrying(date time.Time) string {
	const startYear, startMonth, startDay = 1990, time.January, 1
	startDate := time.Date(startYear, startMonth, startDay, 0, 0, 0, 0, time.Local)
	daysPassed := date.Sub(startDate).Hours() / 24 // 计算经过的天数
	fmt.Println("经过的天数", daysPassed)
	// 每5天为一个周期，前3天打鱼，后2天晒网
	cycleDay := int(daysPassed) % 5 //将经过的天数的类型float64换成int然后取余
	if cycleDay < 3 {
		return "打鱼"
	}
	return "晒网"
}

func main() {
	var year, month, day int
	fmt.Print("请输入年份: ")
	fmt.Scan(&year)
	fmt.Print("请输入月份: ")
	fmt.Scan(&month)
	fmt.Print("请输入日期: ")
	fmt.Scan(&day)

	dateToCheck := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	result := isFishingOrDrying(dateToCheck)
	fmt.Printf("%d年%d月%d日是：%s\n", year, month, day, result)
}
