package main
import "fmt"
func main(){
	var property int
	var height int
	var face bool
	fmt.Println("财产,身高多少厘米,好看不选择输入true和false,中间用空格隔开")
	fmt.Scanln(&property, &height, &face)
	if  property >= 1000000 && height >= 180 && face == true {
		fmt.Println("结婚")
	} else if property >= 1000000 || height >= 180 ||  face == true {
		fmt.Println("结婚吧，比上不足比下有余")
	} else if property < 1000000 && height < 180 &&  face == false {
		fmt.Println("滚")
	} 
}