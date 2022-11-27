package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 电脑
type Computer struct {
}

// var usber Usber= camera
func (c Computer) work(u Usber) {
	u.start()
	u.stop()

}

// 手机
type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println(p.name, "开机")

}
func (p Phone) stop() {
	fmt.Println(p.name, "关机")

}

// 照相机
type Camera struct {
}

func (c Camera) start() {
	fmt.Println("照相启动")

}
func (c Camera) stop() {
	fmt.Println("照相关机")

}

func main() {
	var computer = Computer{}
	var phone = Phone{
		name: "小米",
	}
	computer.work(phone)

	var camera = Camera{}
	computer.work(camera)

}
