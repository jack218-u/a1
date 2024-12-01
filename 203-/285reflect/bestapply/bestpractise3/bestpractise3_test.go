/*
代码解释：
定义了两个函数 call1 和 call2，分别接受两个整数参数和三个参数（两个整数和一个字符串）。
定义了一个 bridge 函数，用于通过反射调用传入的函数。
在 TestReflectFunc 函数中，分别调用了 bridge 函数两次，一次传递 call1 和两个整数参数，
另一次传递 call2 和两个整数及一个字符串参数。
*/

package test

import (
	"reflect"
	"testing"
)

// 代码的文件名以_test.go名称结尾,用go test -v加文件名测试
func TestReflectFunc(t *testing.T) {
	call1 := func(v1, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value   // 用于存储反射值
		inValue  []reflect.Value // 用于存储函数参数的反射值
		n        int             // 参数数量
	)
	// bridge 函数用于通过反射调用任意函数
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)                      // 获取参数数量
		inValue = make([]reflect.Value, n) // 初始化 inValue 切片，用于存储参数的反射值
		for i := 0; i < n; i++ {           // 将参数转换为反射值并存储到 inValue 中
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call) // 将传入的函数转换为反射值
		function.Call(inValue)           // 使用反射调用函数
	}
	bridge(call1, 1, 2)          // 调用 call1 函数，传递两个整数参数
	bridge(call2, 1, 2, "test2") // 调用 call2 函数，传递两个整数和一个字符串参数
}

/*
详细解释
定义函数：call1 和 call2 分别接受不同数量和类型的参数，并使用 t.Log 打印这些参数。
定义 bridge 函数：bridge 函数接受一个函数和任意数量的参数，通过反射机制调用该函数。
初始化参数：将传入的参数转换为 reflect.Value 类型的切片。
反射调用：使用 reflect.ValueOf 获取函数的反射值，并调用 Call 方法执行该函数。
调用 bridge 函数：在 TestReflectFunc 中，分别调用 bridge 函数两次，一次传递 call1
和两个整数参数，另一次传递 call2 和两个整数及一个字符串参数。
*/
/*
控制流图
flowchart TD
    Start[开始] --> DefineFunctions[定义 call1 和 call2 函数]
    DefineFunctions --> DefineBridge[定义 bridge 函数]
    DefineBridge --> CallBridge1[调用 bridge(call1, 1, 2)]
    CallBridge1 --> CallBridge2[调用 bridge(call2, 1, 2, "test2")]
    CallBridge2 --> End[结束]

    subgraph BridgeFunction[bridge 函数内部逻辑]
        InitArgs[初始化参数]
        ReflectCall[通过反射调用函数]
    end

    CallBridge1 -->|进入 bridge 函数| InitArgs
    InitArgs --> ReflectCall
    ReflectCall -->|返回| CallBridge1

    CallBridge2 -->|进入 bridge 函数| InitArgs
    InitArgs --> ReflectCall
    ReflectCall -->|返回| CallBridge2
*/
