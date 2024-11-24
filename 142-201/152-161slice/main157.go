package main

import "fmt"

func main() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//切片需要注意事项和细节说明
	//1. 切片初始化时候var slice := arr[startIndex:endIndex]
	//说明，从arr数组下标为startIndex,取到下标为endIndex的元素(不包含arr[endIndex])
	//2. 切片初始化时，仍然不能越界。范围在[0-len(arr]之间，但是可以动态增长
	//2.1 var slice = arr[0:end]可以简写为：var slice = arr[:end]
	//2.2 var slice = arr[start:len(arr)]可以简写为：var slice = arr[start:]
	//2.3 var slice = arr[0:len(arr)]可以简写：var slice = arr[:]
	// slice := arr[1:4] //20,30,40
	// slice := arr[0:len(arr)] //即切片把数组所有元素都取到了,也等价于arr[:]
	slice := arr[:]   // 10, 20, 30, 40, 50
	slice2 := arr[:3] //10, 20, 30
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
	fmt.Println()
	//使用for-range方式遍历
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
	//3. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
	//4. 切片定义完成后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
	//5. 切片可以继续切片,引用规则类似切片应用数组，取值范围不包含[]中的最后一位
	//5.1 如果被应用的切片发生变化，引用的切片也会发生变化
	slice3 := slice2[1:2] // 20
	slice3[0] = 100       //因为arr,slice ,slice2,slice3 都是指向的数据空间都是同一个，因此slice3[0]变化，其他都会变化
	fmt.Println("slice3=", slice3)
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)
	// 切片扩容
	//6. append内置函数，可以对切片进行动态追加，后面直接跟具体数组，或者跟切片，如果跟切片，后面需要有3个.不能少
	var slice4 []int = []int{100, 200, 300}
	slice5 := append(slice4, 400, 500, 600) //通过append直接给slice4追加具体的元素,然后传递给slice5
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)
	slice5 = append(slice5, 700, 800) //通过append直接给slice5追加元素，改变slice5的值
	fmt.Println("slice5=", slice5)
	slice4 = append(slice4, slice4...) //通过append将切片slice4追加给slice4
	fmt.Println("slice4=", slice4)
	slice5 = append(slice5, slice4...) //通过append将切片slice5追加给slice4
	fmt.Println("slice5=", slice5)
	//append操作的底层原理分析：
	//1. append操作本质就是对数组扩容
	//2. go底层会创建一下新的数组newArr(安装扩容后的大小)
	//3. 将slice原来包含的元素拷贝到新的数组newArr
	//4. slice重新引用到newArr
	//5. 注意newArr是底层来维护的，程序员不可见

	//切片的拷贝操作,copy(slice1,slice2),就是slice2对slice的覆盖，从下标0开始，不足的位置不动，超出的位置省略
	//比如slice1=[1 2 3 4 5 0 0 0 0 0],如果slice2=[123],则结果copy(slice1,slice2)结果不变，如果slice2=[1 2 3 4 5 6]
	//结果为[1 2 3 4 5 6 0 0 0 0],如果slice2=[1 2 3 4 5 6 7 8 9 10 11]，结果为[1 2 3 4 5 6 7 8 9 10],因为slice2的长度为10
	//注意：拷贝操作必须都是切片
	var a []int = []int{1, 2, 3, 4, 5, 6}
	var slice6 = make([]int, 10) //容量是可选值 元素为默认的10个0
	fmt.Println("slice6=", slice6)
	copy(slice6, a) //将a追加给slice6以后，占据前6个位置，后面4个继续为0
	fmt.Println("slice6=", slice6)

	var slice7 []int = []int{1, 2, 3, 4, 5}
	slice8 := make([]int, 10)
	copy(slice8, slice7)
	fmt.Println("slice7=", slice7)
	fmt.Println("slice8=", slice8)
	//修改被添加的切片，对copy的切片无影响，copy的两个slice是独立的空间
	slice7[1] = 20 //修改slice7的第二个下标，对slice8无影响
	fmt.Println("slice7=", slice7)
	fmt.Println("slice8=", slice8)
	// 切片拷贝操作总结：
	//1. copy(slice1,slice2)参数的类型需要是切片
	//2. 按照上面代码来看，copy的两个slice是独立的空间，相互不影响，例如slice7[0]修改成999，slice8[0]仍然是1
	slice9 := []int{1, 2, 3}
	slice8 = append(slice8, slice9...) //在slice8后面增加3个位置，长度变成13
	fmt.Println("slice8扩容-append后增加3个长度= ", slice8)
	slice10 := []int{1, 2, 3}
	copy(slice8, slice10)
	fmt.Println("copy拷贝slice10到slice8后无变化=", slice8)
	slice11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	copy(slice8, slice11)
	fmt.Println("copy拷贝slice11到slice8后的变化slice11覆盖原先的位置,多余的两\n个元素14和15由于长度不够省略,结果为=", slice8)
	fmt.Printf("slice8的数据类型=%T", slice8) //数据类型是切片
}
