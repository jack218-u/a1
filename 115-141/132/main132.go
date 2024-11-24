package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1. len()统计字符串长度，按照字节返回，不是按照字符个数返回
	str := "hello" //golang的编码统一为utf-8，ascii的字符（字母和数字）占一个字节，汉字占3个字节
	fmt.Printf("str的类型=%T, str的长度等于%v个字节\n", str, len(str))
	str1 := "hello北京"
	fmt.Printf("str1的长度等于%v个字节\n", len(str1))
	//2. 字符串遍历
	for i := 0; i < len(str1); i++ {
		// fmt.Println(i, "str1的字符=", str1[i])
		// fmt.Printf("%v str1的字符=%v\n", i, str1[i])
		fmt.Printf("%v str1的字符=%c\n", i, str1[i])
	}
	str2 := "hello北京"
	r := []rune(str2) //字符串遍历，处理有中文的问题： r := []rune(str) 转成rune的切片
	for i := 0; i < len(r); i++ {
		fmt.Printf("%v r的字符=%c\n", i, r[i])
	}
	//3. 字符串转整数：n,err := strconv.Atoi("123")
	n, err := strconv.Atoi("123") //如果这里123替换成hello，则会报错，此处只能输入数字字符串，不能输入字母汉字类型的字符串
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果", n)
	}
	//4. 整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str的类型=%T\n", str, str)
	//5. 字符串转byte： var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	//6. []byte 转 字符串 : str = string([]byte{97,98,99})
	str = string([]byte{97, 98, 99}) //分别对应字母a b c
	fmt.Printf("str=%v,str的类型=%T\n", str, str)
	//7.    10进制转2进制 8进制 16进制 FormatInt(i int, base int) string
	str = strconv.FormatInt(123, 2)
	fmt.Printf("str的数据类型是=%T, 123对应的二进制=%v\n", str, str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("str的数据类型是=%T, 123对应的十六进制=%v\n", str, str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("str的数据类型是=%T, 123对应的八进制=%v\n", str, str)
	//8. 查找子串是否在指定的字符串中：strings.Contains(s,substr string) bool
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b的数据类型=%T, b=%v\n", b, b)
	c := strings.Contains("seafood", "many")
	fmt.Printf("c的数据类型=%T, c=%v\n", c, c)
	//9. 统计一个字符串有几个指定的子串：strings.Count(s,sep string) int
	num := strings.Count("ceheese", "e")
	fmt.Printf("num的数据类型=%T, num=%v\n", num, num)
	//10. 不区分大小写的字符串比较（==是区分字母大小写的）：fmt.Pringln(strings.EqualFold("abc","Abc"))//true
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)
	fmt.Println("结果", "abc" == "Abc") //false
	//11. 返回子串在字符串第一次出现的index值，如果没有返回 -1 : strings.Index("NLT_abc","abc")
	index := strings.Index("NLT_abc", "abc") //计算index值应该从0开始，abc出现在NLT_abc的第4个字符位置
	fmt.Printf("index=%v\n", index)
	index1 := strings.Index("NLT_abc", "hello") // NLT_abc中没有hello，返回-1
	fmt.Printf("index1=%v\n", index1)
	index2 := strings.Index("NLT_abcabcabc", "abc") // 有多个abc，只找第一次出现的index值
	fmt.Printf("index2=%v\n", index2)
	//12. 返回子串在字符串最后一次出现的index，如果没有返回-1 ：strings.LashtIndex("go golang","go")
	index = strings.LastIndex("go golang", "go") //求go在前面"go golang"中最后出现的位置，下标位置以0开头计数
	fmt.Printf("index=%v\n", index)
	//13. 将指定的子串替换成 另外几个子串：strings.Replace("go go hello", "go", "北京",n)
	//n可以指定希望替换几个，如果n= -1表示全部替换，n=1表示替换从前往后的第一个
	str = strings.Replace("go go go hello", "go", "北京", 2)
	fmt.Printf("str=%v\n", str)
	str99 := "go go go hello"
	str98 := strings.Replace(str99, "go", "北京", 2) //strings.Replace函数中也可以用变量表示参数，替换完成后str99本身没有变化
	fmt.Printf("str98=%v,str99=%v\n", str98, str99)
	//14. 按照指定的某个字符进行分割，为分割表示，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
	strArr := strings.Split("hello,world,ok", ",") //其中"hello,world,ok"也可以是变量，传递以后变量值不变，因为是值拷贝
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr) //变成数组
	//15. 将字符串的字母进行大小写的转换：strings.ToLower("Go") //go  strings.ToUpper("Go") //GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)
	//16. 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn ")
	str = strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str=%q\n", str) //%q表示用双引号引起来的%v，看的更清楚
	//17. 将字符串左右两边指定的字符去掉：strings.Trim("! hello!"," !")
	str = strings.Trim("! he!llo! ", " !") //去掉左右两边的!和空格
	fmt.Printf("str=%q\n", str)            //%q表示用双引号引起来的%v，看的更清楚
	//18. 将字符串左边指定字符去掉：strings.TrimLeft("! hello!"," !") 将"! hello!"左边的!和空格去掉
	str = strings.TrimLeft("! he!llo! ", " !") //去掉左边的!和空格
	fmt.Printf("str=%q\n", str)                //%q表示用双引号引起来的%v，看的更清楚
	//19. 将字符串左边指定字符去掉：strings.TrimRight("! hello! "," !") 将"! hello!"左边的!和空格去掉
	str = strings.TrimRight("! he!llo! ", " !") //去掉左边的!和空格
	fmt.Printf("str=%q\n", str)                 //%q表示用双引号引起来的%v，看的更清楚
	//20. 判断字符串是否以某个字符串开头：strings.HasPrefix("ftp://192.168.0.1","ftp") //true
	b = strings.HasPrefix("ftp://192.168.0.1", "ftp")
	fmt.Printf("b=%v\n", b)
	//21. 判断字符串是否以某个字符串结束：strings.HasSuffix("NLT_abc.jpg", "abc") //false
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("b=%v\n", b)
}
