package service

import "abc/203-/customerManager/model"

//该CustomerService,完成对Customer的操作,包括增删改查
type Service struct {
	c1 []model.Customer
	//声明一个字段,表示当前切片含有多少给客户
	//该字段后面,还可以作为新客户的id+1
	num int
}

//编写一个方法,可以返回*CustomerService
func NewService() *Service { //函数名称首字母大写，包才能给外部引用
	//为了能够看到有客户在切片中，我们初始化一个客户
	s1 := &Service{}
	s1.num = 1
	c2 := model.NewCustomer(1, "张三", "男", 20, "13888888888", "zs@gmail.com") //此处是一个方法取值，应该用()括起来
	s1.c1 = append(s1.c1, c2)
	return s1
}

//返回客户切片
func (p *Service) List() []model.Customer {
	return p.c1
}

//添加客户到Servie切片
//这里必须要用指针类型*Service
func (p *Service) Add(c3 model.Customer) bool { //重点！！！，这里用的是指针类型，如果是值类型，新增加客户后，老客户就看不到了
	//我们确定一个分配id的规则，就是添加的顺序
	p.num++       //重点
	c3.Id = p.num //重点
	p.c1 = append(p.c1, c3)
	return true
}

//根据id查找客户在切片中对应下标,如果没有该客户,返回-1
func (p *Service) FindById(id int) int {
	index := -1
	//遍历model.Customer切片
	for i := 0; i < len(p.c1); i++ {
		if p.c1[i].Id == id {
			index = i
			break
		}
	}
	return index
}

//根据id删除客户(从切片中删除)
func (p *Service) DeleteById(id int) bool {
	index := p.FindById(id)
	//如果index == -1,说明没有找到该客户
	if index == -1 {
		return false
	}
	//从切片中删除一个元素
	p.c1 = append(p.c1[:index], p.c1[index+1:]...)
	//切片是左边右开,p.c1[:index)表示取到index-1,p.c1[index+1:)表示从index+1开始取到切片末尾]
	//不包括p.c1[index]
	return true
}

//根据id修改用户信息
func (p *Service) Update(id int, name string, gender string, age int, phone string, email string) bool {
	index := p.FindById(id)
	//如果index == -1,说明没有找到该客户
	if index == -1 {
		return false
	}
	//更新客户信息
	p.c1[index].Name = name
	p.c1[index].Gender = gender
	p.c1[index].Age = age
	p.c1[index].Phone = phone
	p.c1[index].Email = email
	return true
}
