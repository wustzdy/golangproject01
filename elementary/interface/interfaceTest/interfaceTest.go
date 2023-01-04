package main

import "fmt"

// 实现接口
type Usber interface {
	start()
	stop()
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

// 手机结构体
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

// 照相机结构体
type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机启动")
}
func (c Camera) stop() {
	fmt.Println("相机关机")
}

func main() {
	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	computer.work(phone)
	/*
		小米 启动
		小米 关机
	*/

}
