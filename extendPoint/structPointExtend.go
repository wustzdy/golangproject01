package main

import "fmt"

// 父类结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.Name)
}

// 子结构体
type Dog struct {
	Age     int
	*Animal //结构体嵌套 继承
}

func (d Dog) wangwang() {
	fmt.Printf("%v 在望望\n", d.Name)
}

func main() {
	d := Dog{
		Age: 12,
		Animal: &Animal{
			Name: "阿强",
		},
	}
	d.run()
	d.wangwang()
}
