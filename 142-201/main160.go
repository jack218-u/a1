package main

import "fmt"

func main() {
	//1. string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@atguigu"
	//使用切片获取到atguigu
	slice := str[3:]
	fmt.Println("slice=", slice)
	// fmt.Printf("%T\n", slice)
	//2. string和切片在内存的形式，以"abcd"画出内存示意图
	//3. string不能通过str[0]='z'方式来修改字符串的内容
	// str[0] = z 这样子修改会报错，编译不会通过，原因是string是不可变的
	//4. 如果需要修改字符串，可以先将string->[]byte或者[]rune->修改->重写成string
	arr1 := []byte(str)
	arr1[0] = 'z'
	fmt.Println(arr1)
	str = string(arr1)
	fmt.Println(str)
	//细节，转成[]byte后，可以处理英文和数字，但是不能处理中文
	//原因是，[]byte是按字节来处理的，但是汉字是3个字节，因此会出现乱码
	//解决方法是，将 string转成[]rune即可，因为[]rune是字符处理，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	fmt.Println(arr2)
	str = string(arr2)
	fmt.Println(str)

}
