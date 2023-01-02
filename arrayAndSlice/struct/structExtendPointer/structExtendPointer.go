package main

import "fmt"

// 结构体继承

// 父结构体
type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.name)
}

// 子结构体
type Dog struct {
	Age     string
	*Animal //发现Dog结构体没有name属性，我们可以继承父类Animal结构体 叫结构体继承 也叫嵌套
}

func (d Dog) wang() {
	fmt.Printf("%v 在汪汪\n", d.name)
}

func main() {
	var d = &Dog{
		Age: "20",
		Animal: &Animal{
			name: "汪曾祺",
		},
	}
	d.run() //dog继承了Animal的run方法
	d.wang()
	/*
		汪曾祺 在运动
		汪曾祺 在汪汪
	*/

}
