package main
import "fmt"
func main(){
	//基本数据类型在内存布局
	var i int = 10
	//i的地址是什么，&i
	fmt.Println("i的地址=",&i)
	//下面的var ptr *int = &i
	//1. ptr是指针变量
	//2. ptr的类型 *int，*int表示指向int数据的指针数据
	//3. ptr本身的值/内存空间地址&i,或者&(ptr)
	//4. ptr对外显示的值/指针指向的值是*ptr
	var ptr *int = &i
	fmt.Printf("*ptr的值=%v,ptr的数据类型是:%T,ptr的值=%v,&ptr的值=%v\n",*ptr,ptr,ptr,&(ptr))
	var ptr2 **int = &(ptr) //嵌套两层
	fmt.Printf("*ptr2的值=%v,ptr2的数据类型是:%T,ptr2的值=%v,&(ptr2)的值=%v\n",*ptr2,ptr2,ptr2,&(ptr2))
	var ptr3 ***int = &(ptr2) //嵌套三层
	fmt.Printf("*ptr3的值=%v,ptr3的数据类型是:%T,ptr3的值=%v,&(ptr3)的值=%v\n",*ptr3,ptr3,ptr3,&(ptr3))

	var a int = 300
	var b int = 400
	var ptr1 *int = &a //指针prt1指向变量a
	*ptr1 = 100 //指针ptr1修改为100，则a变成100
	ptr1 = &b   //指针ptr1指向变量b
	*ptr1 = 200 //指针ptr1指向200，则b变成200
	fmt.Printf("a=%d,b=%d,*ptr1=%d",a,b,*ptr1)
}
