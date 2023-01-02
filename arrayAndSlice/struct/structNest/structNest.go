package main

import "fmt"

// 结构体嵌套
type User struct {
	UserName string
	PassWord string
	Sex      string
	Age      int
	Address  Address //user结构体嵌套Address结构体
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

	fmt.Printf("%#v", u) //main.User{UserName:"张三", PassWord:"1233", Sex:"", Age:0, Address:main.Address{Name:"李四", Phone:"1223333333", City:"北京"}}
}
