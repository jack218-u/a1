package main

import "fmt"

// 练习题
// 1. 使用map[string]map[string]string
// 2. key表示用户名，是唯一的，不可以重复
// 3. 如果某个用户名存在，就将其密码修改为"888888",如果不存在就增加这个用户信息(包括昵称nickname和密码pwd)
// 4. 编写一个函数modifyUser(user map[string]map[string]string,name string)完成上述功能
// modifyUser函数的所用是判断引入的参数name是否在类型为map的参数user中存在，然后执行判定条件
func modifyUser(users map[string]map[string]string, name string) {
	//判断users中是否有name，可以用查询函数，或者if语句
	//v,ok := users[name] 查询
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//目前没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["nickname"] = "昵称~" + name //示意代码
		users[name]["pwd"] = "888888"
	}
}
func main() {
	users := make(map[string]map[string]string, 10) //假设增加10个
	users["john"] = make(map[string]string, 2)
	users["john"]["nickname"] = "黑熊"
	users["john"]["pwd"] = "999999"
	// user["No.1"] = make(map[string]string)
	// user["No.2"] = make(map[string]string)
	// user["No.1"]["name"] = "九纹龙-史进"
	// user["No.2"]["name"] = "神机军师-朱武"
	// user["No.1"]["pwd"] = "777777"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "john")
	fmt.Println("john使用函数前的密码", users["john"]["pwd"])
	fmt.Println("使用函数后的变化", users)

}
