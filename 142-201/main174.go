package main

import "fmt"

func main() {
	//map的增删改查
	//1. map增加和更新：
	//map["key"]= value //如果key没有，就是增加，如果key存在就是修改&更新
	cities := make(map[string]string)
	cities["No.1"] = "北京"
	cities["No.2"] = "天津"
	cities["No.3"] = "上海"
	fmt.Println(cities)
	cities["No.3"] = "上海~" //因为No.3已经存在，所以下面的语句就是修改
	fmt.Println(cities)
	//2. map删除
	//说明，delete(map,"key"),delete是一个内置函数，如果key存在，就删除key-value，如果key不存在，不操作，但是也不会报错
	delete(cities, "No.1") //删除No.1
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "No.4")
	fmt.Println(cities)
	//map删除细节说明
	//a. 如果我们要删除map的所有key，没有一个专门的方法一次性删除，可以遍历一下key，逐个删除
	//b. 或者map = make(...),make一个新的，让原来的成为垃圾，被gc回收
	cities = make(map[string]string)
	fmt.Println("make重新之后的结果", cities)
	//3. map查找
	names := map[string]string{"No.1": "史进", "No.2": "朱武", "No.3": "陈达", "No.4": "杨春"}
	fmt.Println(names)
	val, ok := names["No.1"]
	if ok {
		fmt.Printf("有No.1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有No.1 key\n")
	}

	//map遍历
	//map遍历只能用for-rang方式，不能用普通的for循环，因为普通for循环按照下标为0 1 2 3...来操作的，map的key不一定是数字
	for k, v := range names {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	//遍历复杂的map(value也是map的类型)
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 3) //这里的make不能少
	stuMap["stu01"]["name"] = "jack"
	stuMap["stu01"]["sex"] = "男"
	stuMap["stu01"]["addr"] = "北京长安街"
	stuMap["stu02"] = make(map[string]string, 3) //这里的make不能少
	stuMap["stu02"]["name"] = "mary"
	stuMap["stu02"]["sex"] = "女"
	stuMap["stu02"]["addr"] = "上海黄浦江"
	fmt.Println(stuMap) //输出学生表信息
	for k1, v1 := range stuMap {
		// fmt.Printf("k=%v", k1) //这里k后面如果多加一个字符则超过一个制表符，map中v2=mary和v2=男这一行就会右边移动制表符
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
	//统计得到map的长度，func len内置函数
	fmt.Println("names 有", len(names), "对 key-value")
	fmt.Println("stuMap 有", len(stuMap), "对 key-value") //只有两对stu01和stu02，后面的key是另一个map
}
