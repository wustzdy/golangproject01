package main

import "fmt"

// 实现接口
type Usber interface {
	start()
	stop()
}

// 如果接口l里面有方法的话，必须要通过结构体或者通过自定义类型来实现这个接口
type Phone struct {
	Name string
}

// 手机Phone 要实现usb接口话，必须得实现ubs接口里面所有的方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机启动")
}
func (c Camera) stop() {
	fmt.Println("相机关机")
}
func (c Camera) run() {
	fmt.Println("run")
}

func main() {
	p := Phone{
		Name: "华为手机",
	}
	//p.start() //华为手机 启动

	var p1 Usber //golang接口就是一个数据类型
	p1 = p       //表示手机实现Usb接口
	p1.start()   //华为手机 启动
	p1.stop()    //华为手机 关机

	c := Camera{}
	var c1 Usber = c //表示手相机实现Usb接口
	c1.start()       //相机启动
	//c1.run//错误的调用方法
	c.run()
}
