/*
反射最佳实践
1.使用反射来遍历结构体的字段,调用结构体的方法,并获取结构体标签的值
比如func (Value)Method函数
func(Value)Call函数,func(Value)CallSlice函数
2.使用反射的方式来获取结构体的tag标签,遍历字段的值,修改字段的值,
调用结构体方法(要求通过传递地址的方式完成)
*/
package main

import (
	"fmt"
	"reflect"
)

// 定义一个Monster结构体
type Monster struct {
	//此处用json标签,获取tag的时候Get也用json
	//同时如果做反序列化的时候也需要导入encoding/json包
	//如果希望支持多种序列化格式,可以在字段上使用多个标签,如下
	Name    string  `json:"name" xml:"name"`
	Age     int     `json:"monster_age"`
	Score   float32 `json:"成绩"`
	Sex     string
	Address string
}

// 方法显示s的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

// 方法返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 方法 接收四个值,给s赋值
func (s *Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是struct,就退出
	if kd != reflect.Struct { //因为是常量所以可以判断比较
		fmt.Println("expect struct")
		return
	}
	num := val.NumField() // 获取结构体有几个字段
	/*
		func (v Value) NumField() int
		返回v持有的结构体类型值的字段数int，如果v的Kind不是Struct会panic
	*/

	fmt.Printf("struct has %d fields\n", num) // 输出结构体的字段数量,Monster中定义几个就输出几个
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field%d:值为=%v\n", i, val.Field(i))
		/*
			func (v Value) Field(i int) Value
			返回结构体的第i个字段（的Value封装）。如果v的Kind不是Struct或i出界会panic
		*/
		//获取到struct标签,注意需要通过reflect.Type来获取tag标签
		tagVal := typ.Field(i).Tag.Get("json") // 获取字段的JSON标签,Tag是字段的标签
		/*
		   		Field(i int) StructField
		    //   返回索引序列指定的嵌套字段的类型，
		    //   等价于用索引中每个值链式调用本方法，如非结构体将会panic
		*/
		/*
			type StructField struct {
				// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
				Name    string
				PkgPath string
				Type      Type      // 字段的类型
				Tag       StructTag // 字段的标签
				Offset    uintptr   // 字段在结构体中的字节偏移量
				Index     []int     // 用于Type.FieldByIndex时的索引切片
				Anonymous bool      // 是否匿名字段
			}
			StructField类型描述结构体中的一个字段的信息。
		*/
		/*
			func (tag StructTag) Get(key string) string
			Get方法返回标签字符串中键key对应的值。如果标签中没有该键，
			会返回""。如果标签不符合标准格式，Get的返回值是不确定的。
		*/
		//如果该字段有tag标签就显示,否则不显示
		if tagVal != "" { // 如果标签不为空
			fmt.Printf("Field%d:tag为=%v\n", i, tagVal)
		}
		/*
		   		func (v Value) Field(i int) Value
		   返回结构体的第i个字段（的Value封装）。如果v的Kind不是Struct或i出界会panic
		*/
		/*
			func (tag StructTag) Get(key string) string
			Get方法返回标签字符串中键key对应的值。如果标签中没有该键，会返回""。
			如果标签不符合标准格式，Get的返回值是不确定的。
		*/
	}
	numOfMethod := val.NumMethod() // 获取结构体的方法数量
	/*
	   	func (v Value) NumMethod() int
	   返回v持有值的方法集的方法数目
	*/
	fmt.Printf("struct has %d methods\n", numOfMethod)
	//var params[]reflect.Value
	val.Method(1).Call(nil) // 调用结构体的第二个方法，不传递参数,
	//其中排序按照函数名字ascii排序,排第二个的Print函数
	/*
		func (v Value) Method(i int) Value
		返回v持有值类型的第i个方法的已绑定（到v的持有值的）状态的函数形式的Value封装。
		返回值调用Call方法时不应包含接收者；返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）。
		如果i出界，或者v的持有值是接口类型的零值（nil），会panic。
	*/
	/*
		func (v Value) Call(in []Value) []Value
		Call方法使用输入的参数in调用v持有的函数。例如，如果len(in) == 3，v.Call(in)代表调用v(in[0],
		in[1], in[2])（其中Value值表示其持有值）。如果v的Kind不是Func会panic。它返回函数所有输出结果的
		Value封装的切片。和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型。
		如果v持有值是可变参数函数，Call方法会自行创建一个代表可变参数的切片，将对应可变参数的值都拷贝到里面。
	*/

	//调用结构体的第一个方法method(0)
	var params []reflect.Value //声明了[]reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 传入的参数是[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果,返回的结果是[]reflect.Value*/
	//因为知道结果是Int类型,这里直接写Int(),如果不是Int就会报错
	/*
		[]reflect.Value 是一个切片，用于存储多个 reflect.Value 对象。
		在反射调用方法时，参数和返回值都可以被包装成 []reflect.Value 切片。
		通过 reflect.Value.Call 方法可以调用方法并传递参数，返回值也是一个 []reflect.Value 切片。
	*/
}
func main() {
	//创建一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
		//此处没有定义Sex字段和值,就是默认为空,输出空格
	}
	//将Monster实例传给TestStruct函数
	TestStruct(a)
}
