package main
import "fmt"
func main() {
	var name string
	for {
		fmt.Println("输入作者姓名或者图书类别查询图书")
		fmt.Scanln(&name)
		if name == "古龙" {
			fmt.Println("天涯明月刀，流星蝴蝶剑，多强剑客无情剑")
		} else if name == "金庸" {
			fmt.Println("射雕英雄传，神雕侠侣，笑傲江湖")
		} else if name == "六大名著" {
			fmt.Println("三国演义，水浒传，西游记，红楼梦，儒林外史，聊斋志异")
		} else if name == "卫子丹" {
			fmt.Println("黄帝内经-灵柩，皇帝内经-素问")
		} else if name == "国外名著" {
			fmt.Println("月亮与六便士，人性的枷锁，霍乱时期的爱情，瓦尔登湖，浮士德")
		} else {
			fmt.Println("查无此书")
		}
	}
}
