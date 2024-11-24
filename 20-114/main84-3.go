package main

import (
    "fmt"
)

func main() {
    var property, height int
    var face bool

    fmt.Println("请输入财产（元）、身高（厘米）、是否帅气（true/false），用空格隔开:")
    
    // 检查输入是否成功读取
    if _, err := fmt.Scan(&property, &height, &face); err != nil {
        fmt.Println("输入格式错误，请按要求输入。")
        return
    }

    switch {
    case property >= 1000000 && height >= 180 && face:
        fmt.Println("我一定要嫁给他")
    case property >= 1000000 || height >= 180 || face:
        fmt.Println("嫁了吧，比上不足比下有余")
    default:
        fmt.Println("滚")
    }
}