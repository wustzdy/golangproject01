package main

import "fmt"

// 接口嵌套  得到一个新的接口
type Ainterface interface {
	SetName(string)
}
type Binterface interface {
	getName() string
}
type Animal interface { //接口的嵌套
	Ainterface
	Binterface
}
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) getName() string {
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
	fmt.Println(d1.getName()) //小花狗狗
}
