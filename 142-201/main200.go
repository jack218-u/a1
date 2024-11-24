// 创建结构体变量时指定字段值
package main

import "fmt"

type Student3 struct {
	Name string
	Age  int
}

func main() {
	//golang在创建结构体实例(变量)时，可以直接指定字段的值。
	//方式1.
	//1.1 var stu1 Student3 = Student3{"tom",10} //直接输入字段值，按照顺序写，第一个值对应第一个字段，
	//第二个值对应第二个字段，其中Student3也可以省略，写成var Student3 =Student3{"tom",10}
	//tom是第一个字段的值，10是第二个字段的值
	//1.2 或者简写为stu2 := Student3{"tom",10} 用类型推导方式
	//1.3 定义的结构体变量的同时，输入字段名称和字段值(这种写法不依赖字段顺序)，参考如下：
	/*
		var stu3 Student3 = Student3{ //
			Name: "mary", //赋值的时候，中间需要有英文冒号，结果要有英文逗号
			Age:  30,
		}
	*/
	// 1.4 上面1.3的简写为(也可以用类型推导)：
	/*
		stu3 := Student3{
			Name: "mary",
			Age:  30,
		}
	*/
	//方式2，返回结构体的指针类型
	//使用指针方式,var stu1 *Student3 = &Student3{"小明", 20},使用指针的其他类型参考方式1
	/*
		方式2.1
			var stu5 *Student3 = &Student3{"小明", 20} //stu5-->指向地址-->然后指向结构体数据空间
			也可以写成var stu5  = &Student3
		方式2.2, 2.1的简写
			stu6 := &Student3{"小明", 20}
		方式2.3 定义的结构体变量的同时，输入字段名称和字段值(这种写法不依赖字段顺序)
			var stu7 *Student3 = &Student3{
				Name: "小明",
				Age:  20,
			}
		方式2.4,方式2.3的简写
			stu8 := &Student3{
				Name: "小明",
				Age:  20,
			}
	*/
	var stu1 Student3 = Student3{"小明", 20}
	stu2 := Student3{"小明", 20}
	var stu3 Student3 = Student3{
		Name: "小明",
		Age:  20,
	}
	stu4 := Student3{
		Name: "小明",
		Age:  20,
	}
	fmt.Println(stu1, stu2, stu3, stu4)
	var stu5 *Student3 = &Student3{"小王", 29} //也可以省略*Student3,写成var stu5  = &Student3
	stu6 := &Student3{"小王", 39}
	var stu7 *Student3 = &Student3{
		Name: "小李",
		Age:  49,
	}
	stu8 := &Student3{
		Name: "小李",
		Age:  59,
	}
	fmt.Println(stu5, stu6, stu7, stu8)
	fmt.Println(*stu5, *stu6, *stu7, *stu8) //用取值运算符*把地址符号拿掉
	fmt.Printf("stu5值=%p stu6值=%p stu7值=%p stu8值=%p\n", stu5, stu6, stu7, stu8)
	//stu5--8是指针地址，输出结果真实结果需要用%p
	fmt.Printf("stu5本身地址=%p stu6本身地址=%p stu7本身地址=%p stu8本身地址=%p", &stu5, &stu6, &stu7, &stu8)

}
