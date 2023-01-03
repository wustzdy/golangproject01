package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Sex    string
	height int
}

// 结构体指针方法-指针类型的接受者
func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年纪:%v\n", p.Name, p.Age) //姓名:张三 年纪:23
}

func main() {
	var p1 = Person{
		Name: "张三",
		Age:  23,
		Sex:  "女",
	}
	p1.printInfo() // 姓名:张三 年纪:23
	p1.SetInfo("李四", 34)
	p1.printInfo() //姓名:李四 年纪:34

}
