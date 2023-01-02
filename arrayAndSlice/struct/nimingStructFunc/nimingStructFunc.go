package main

import "fmt"

// 匿名结构体
type Person struct {
	string
	int
}

func main() {
	p := Person{
		"张三",
		20,
	}
	fmt.Println("p:", p) //p: {张三 20}
}
