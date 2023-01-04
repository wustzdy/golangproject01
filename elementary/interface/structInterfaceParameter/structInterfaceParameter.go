package main

import "fmt"

// 一个结构体实现多个接口
type Animal1 interface {
	SetName(string)
}
type Animal2 interface {
	GetName() string
}

// 狗
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d *Dog) GetName() string {
	return d.Name
}

func main() {
	DogTest()
}

func DogTest() {
	//Dog实现Animal接口
	var d = &Dog{
		Name: "小黑",
	}
	var d1 Animal1 = d //让dog 实现Animal1这个接口
	var d2 Animal2 = d //让dog 实现Animal2这个接口
	d1.SetName("小花狗狗")
	fmt.Println(d2.GetName()) //小花狗狗

}
