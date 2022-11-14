package main

import "fmt"

type Animal interface {
	eat()
	run()
}
type Dog struct {
	name string
}
type Cat struct {
	name string
	sex  bool
}

func (c Cat) eat() {
	fmt.Println("cat开始吃", c)
}
func (c Cat) run() {
	fmt.Println("cat开始跑", c)
}

func (d Dog) run() {
	fmt.Println("dog开始跑----", d)
}

func main() {
	var a Animal
	a = Cat{
		"小猪", true,
	}
	a.run()
	a.eat()
}
