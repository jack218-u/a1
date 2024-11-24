package main

import "fmt"

func main() {
	// for index, value := range array01 {

	// }
	//1. 第一个返回值index是数组的下标
	//2. 第二个value是该下标位置的值
	//3. 他们都是仅在for循环内部可见的局部变量
	//4. 遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线
	//5. index和value的名称不是固定的，可以自行指定

	heros := [...]string{"宋江", "卢俊义", "吴用"}
	for i, v := range heros {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
	for _, v := range heros { // _表示不关心下标，只关心值
		fmt.Printf("v=%v ", v)
	}
}
