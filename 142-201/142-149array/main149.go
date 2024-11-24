package main

import (
	"fmt"

	"math/rand"
)

func main() {
	// 1)创建一个byte类型的26个元素数组，分别放置A-Z
	// 2）使用for循环访问所有元素并打印出来，提示：字符数据运算'A'+1 -> 'B'
	//思路
	//1.声明一个数组，var myChars [26]byte
	//2. 使用for循环，利用 字符可以进行运算的特点来赋值 'A'+1 -> 'B'
	//3. 使用for打印即可
	//代码：
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) //注意需要将转换成byte,下标为0是是自己即A，后面为1--25，变成B--Z
		fmt.Printf("%c ", myChars[i])
	}
	//求出一个数组的最大值，并得到对应的下标
	//思路
	//1. 声明一个数组var intArr[5] = [...]int {1,-1,9,90,11}
	//2. 假定第一个元素是最大值，下标0
	//3. 然后从第二个元素开始循环比较，如果发现有更大，则交换
	fmt.Println()
	var intArr [5]int = [...]int{100, -1, 9, 90, 6000}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 1; i < len(intArr); i++ { //因为假设的最大值在下标0，所以可以从1开始遍历
		//然后从第二个元素开始循环比较，如果发现有更大，则交换
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)
	//请求出一个数组的和以及平均值，用for-rang
	//思路
	//1.声明一个数组， var intArr [5]int = [...]int {1,-1,9,90,11}
	//2.求出和sum
	//3.求出平均值
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	//如何让平均值保留到小鼠
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr2)))

	//要求：随机生成5个数，并将其反转打印
	//思路
	//1. 随机生成5个数，用rand.Intn()函数
	//2. 当得到随机数后，放到数组中，比如int数组中
	//3.反转打印，交换的次数，是数组个数/2，倒数第一和第一交换，倒数第二个和第二个交换

	var intArr3 [5]int
	len := len(intArr3) //提前对数组长度做赋给一个变量
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100) //新版本每次生成的随机数不一样，不需要种子时间seed函数
	}
	fmt.Println("交换前", intArr3)
	//反转打印，交换的次数是 len / 2 ,len是整数，结果省略小数部分
	//倒数第一个和第一个元素交换，倒数第二个和第二个元素交换
	tmp := 0 //做一个临时变量
	for i := 0; i < len/2; i++ {
		tmp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = tmp
	}
	fmt.Println("交换后", intArr3)
	var arr19 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr19=", arr19)

}
