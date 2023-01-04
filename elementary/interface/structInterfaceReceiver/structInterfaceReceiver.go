package main

import "fmt"

// 结构体值接受者实现接口
type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() { //值接受者
	fmt.Println(p.Name, "启动")
}
func (p Phone) Stop() {
	fmt.Printf(p.Name, "关机")
}

func main() {
	//结构体值接收者实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	var p1 = Phone{
		Name: "小米手机",
	}
	var p2 Usb = p1 //表示让Phone实现Usb接口
	p2.Start()      //小米手机 启动

	var p3 = &Phone{
		Name: "华为手机",
	}
	var p4 Usb = p3 //表示让Phone实现Usb接口
	p4.Start()      //华为手机 启动

}
