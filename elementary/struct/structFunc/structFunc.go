package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Sex    string
	height int
}

// 结构体方法
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年纪:%v\n", p.Name, p.Age) //姓名:张三 年纪:23
}

func main() {
	var p1 = Person{
		Name: "张三",
		Age:  23,
		Sex:  "女",
	}
	p1.printInfo() //p1实例

	var p2 = Person{
		Name: "李四",
		Age:  89,
		Sex:  "男",
	}
	p2.printInfo() //p2实例

}
