package main

import "fmt"

// /声明结构体
// 字段名和函数名一样，首字母大写，在包外可以访问
type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s *Student) setStudent(name string) {
	s.Name = name
}

func main() {
	s1 := Student{
		Name:  "李四",
		Age:   12,
		Score: 89.5,
	}
	(&s1).setStudent("张三")
	fmt.Println("s1:", s1)
}
