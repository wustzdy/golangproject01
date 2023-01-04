package main

import "fmt"

// 接口嵌套  得到一个新的接口
type AInterface interface {
	SetName(string)
}
type BInterface interface {
	GetName() string
}
type Animal interface { //接口的嵌套
	AInterface
	BInterface
}
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}

func main() {
	interfaceNest()
}

func interfaceNest() {
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animal = d
	d1.SetName("小花狗狗")
	fmt.Println(d1.GetName()) //小花狗狗
}
