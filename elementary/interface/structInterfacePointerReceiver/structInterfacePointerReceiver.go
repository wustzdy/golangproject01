package main

import "fmt"

// 指针接收者：如果结构体的方法是指针接收者，那么实例化后的结构体指针类型都可以赋值给变量接口
type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p *Phone) Start() { //指针接收者
	fmt.Println(p.Name, "启动")
}
func (p *Phone) Stop() { //指针接收者
	fmt.Printf(p.Name, "关机")
}

func main() {
	//结构体值接收者实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	var phone1 = &Phone{
		Name: "小米手机",
	}
	var u Usb = phone1 //表示让Phone实现Usb接口
	u.Start()          //小米手机 启动
}
