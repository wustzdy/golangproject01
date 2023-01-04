package main

import "fmt"

// 接口参数实现
type Animal interface {
	SetName(string)
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

// 猫
type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c *Cat) GetName() string {
	return c.Name
}
func main() {
	DogTest()
	CatTest()

}

func CatTest() {
	fmt.Println("----------------------------")
	//Cat实现Animal接口
	var c = &Cat{
		Name: "小猫",
	}
	var animal1 Animal = c
	fmt.Println(animal1.GetName()) //小猫

	//setName
	animal1.SetName("小花猫")
	fmt.Println(animal1.GetName()) //小花猫
}

func DogTest() {
	//Dog实现Animal接口
	var d = &Dog{
		Name: "小黑",
	}
	var animal Animal = d
	fmt.Println(animal.GetName()) //小黑

	//setName
	animal.SetName("小花")
	fmt.Println(animal.GetName()) //小花
}
