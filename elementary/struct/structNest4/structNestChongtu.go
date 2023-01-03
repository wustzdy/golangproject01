package main

import "fmt"

// 结构体嵌套-命名冲突赋值问题
type User struct {
	UserName string
	PassWord string
	Address  //user结构体嵌套Address结构体
	Email    Email
}
type Address struct {
	Name    string
	Phone   string
	City    string
	AddTime string
}
type Email struct {
	Account string
	AddTime string
}

func main() {
	var u User
	u.UserName = "张三"
	u.PassWord = "1233"

	u.Address.Name = "李四"
	u.Address.Phone = "1223333333"
	u.Address.City = "北京"

	fmt.Printf("%#v\n", u) //main.User{UserName:"张三", PassWord:"1233", Sex:"", Age:0, Address:main.Address{Name:"李四", Phone:"1223333333", City:"北京"}}

	u.Address.AddTime = "2020--9999" //更改的是Address结构体的Address
	u.Email.AddTime = "2020--1112"   //更改的是Email结构体的Address
	fmt.Printf("%#v\n", u)
}
