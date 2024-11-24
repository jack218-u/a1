package main
import (
	"fmt"
	"unsafe"
)
func main(){
	var b = false
	fmt.Println("b=",b)
	fmt.Printf("b占用的字节数:%d \n",unsafe.Sizeof(b))
	fmt.Printf("b对应的数据类型是:%T",b)
}