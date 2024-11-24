//工厂模式
/*
一个结构体声明如下：
package model
type Student struct {
	Name string...
}
因为这里的Student的首字母S是大写的,如果我们想在其他包创建Student的实例,比如main包,引入model包后,就可以直接创建Student结构体的变量(实例).
但是问题来了,如果首字母是小写的,比如是type student struct {...}就不行了,如何解决--->工厂模式解决
*/
//使用工厂模式来解决问题
//使用工厂模式实现跨包创建结构体实例(变量)的案例
//1.如果model包的结构体变量首字母大写,引入后，直接使用没问题
//2.如果model包的结构体变量首字母小写，引入后，不能直接使用，可以用工厂模式解决，案例见factory/model