package main
import "fmt"
func main(){
	//先判断再执行
		var i int = 1
		for {
			if i > 10 {
			break
		}
		fmt.Println("helloworld",i)
		i++
	}
	//先执行再判断
	var j int = 1
	for {
		fmt.Println("你好",j)
		j++
		if j > 10{
			break
		}
	}
}