package main

import (
	"fmt"
	"reflect"
)

// 基本数据类型演示反射
func reflectTest01(b interface{}) {
	//通过反射获取到传入的变量type(类型),kind(类别),值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b) //TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil
	fmt.Printf("rTyp=%v,  rType的类型是=%T\n", rTyp, rTyp)
	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	// n2 := 200 + rVal 此操作是错误的,需要加reflect包中的Int函数
	n2 := 2 + rVal.Int() //如果传入的值不是Int类型的就需要断言
	// n3 := 2 + rVal.Float() //此时编译不会报错,但是执行输出会报错
	//panic: reflect: call of reflect.Value.Float on int Value错误
	// fmt.Println(n3)
	fmt.Printf("rVal的值=%v,  rVal的类型是=%T\n", rVal, rVal)
	fmt.Println("n2的值=", n2)
	//下面将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv的值=%v,iv的类型=%T\n", iV, iV) //此时的类型虽然变成int,但是还需要断言
	// fmt.Println(iV + 100)此时仍然不能和其他数据类型相加,因为缺少数据类型,需要做断言
	//将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2的值=%v num2的类型=%T\n", num2, num2)
	fmt.Println()
	/*
		以上步骤总结:
		1.先将main主函数中的变量int类型的变量num取值,值=100,reflect.Value函数操作赋值给变量rVal
		2.将rVal转换成空接口赋值给变量iV,此时运行时能判断类型,但是编译器不能判断Interface()函数操作
		3.对空接口类型数据断言转换成原有int类型
	*/
}

type Student struct {
	Name string
	Age  int
}

// 演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取到传入的变量type(类型),kind(类别),值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b) //TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil
	fmt.Printf("rTyp=%v,  rType的类型是=%T\n", rTyp, rTyp)
	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal的值=%v,  rVal的类型是=%T\n", rVal, rVal)
	//下面将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv的值=%v iv的类型=%T\n", iV, iV)
	//此时类型虽然是Student结构体类型,但是无法单独读取其中字段的值,比如iV.Name
	//反射的本质是运行时能确定结构体的类型和值,但是编译阶段没法确定,只能断言
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
	//3.获取 变量对应的Kind(类别),kind是比较大的范畴,比type范畴大
	//方式a. rVal.Kind()获取
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	//方式b. rType.Kind()获取
	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)
}

/*
反射需要注意的事项和细节说明
1.reflect.Value.kind,获取的变量类别,返回的是一个常量
比如kind返回v持有的值的分类,如果v是value零值,返回值是invaild
Kind代表Type类型值表示的具体分类.零值表示非法分类
2.Type是类型,Kind是类别,Type和Kind可能是相同的,也可能是不同的
比如: var num int = 10, num的type是int,kind也是int
比如: var stu Student,stu的type是Student,kind是struct
3.通过反射可以在让变量在interface{}和Reflect.Value之间相互转换
4.使用反射的方式来获取变量的值(并返回相应的类型),要求数据类型匹配,比如x是int,
那么就应该使用reflect.Value(x).Int(),而不能使用其他,比如17行代码,结构体不能使用此方法,需要先断言
5.通过反射来修改变量,注意当使用SetXxx方法来设置需要通过对应的指针类型来完成,这样才能改变传入的变量的值,
同时需要使用到reflect.Value.Elem()方法
*/
func main() {
	//演示对(基本数据类型,interface{}、reflect.Value)进行反射的基本操作
	//1. 定义一个int
	var num int = 100
	reflectTest01(num)
	//2. 定义一个Student实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
	var num2 int = 100
	fn := reflect.ValueOf(&num2) //通过反射获取到num2的指针地址并赋值给fn
	fn.Elem().SetInt(200)        //通过Elem()方法获取到num2的值,然后通过SetInt()方法修改num2的值
	fmt.Printf("%v\n", num2)
	/*
		Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
		如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值
	*/
	// 等价于:
	var num3 int = 100
	var b *int = &num3
	*b = 200
}
