package main

import "fmt"

// 结构体嵌套
type User struct {
	UserName string
	PassWord string
	Sex      string
	Age      int
	Address  //user结构体嵌套Address结构体
}
type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {
	var u User
	u.UserName = "张三"
	u.PassWord = "1233"

	u.Address.Name = "李四"
	u.Address.Phone = "1223333333"
	u.Address.City = "北京"

	fmt.Printf("%#v\n", u) //main.User{UserName:"张三", PassWord:"1233", Sex:"", Age:0, Address:main.Address{Name:"李四", Phone:"1223333333", City:"北京"}}

	//现对Address结构体进行赋值
	u.City = "上海"        //当访问结构体成员时，会先在结构体中查找，如果怕查找不到，再去匿名结构体中查找
	fmt.Printf("%#v", u) //main.User{UserName:"张三", PassWord:"1233", Sex:"", Age:0, Address:main.Address{Name:"李四", Phone:"1223333333", City:"上海"}}
}
