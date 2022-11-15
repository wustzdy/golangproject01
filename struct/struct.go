package main

import "fmt"

// /声明结构体
// 字段名和函数名一样，首字母大写，在包外可以访问
type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	//struct字段访问，使用“点”访问
	//var stu Student
	var stu *Student = new(Student)
	//var stu *Student = &Student{ }
	stu.Name = "yuchao"
	stu.Age = 18
	stu.Score = 99
	fmt.Println(stu) //默认结构 {yuchao 18 99.1}
	//fmt.Printf("name=%s age=%d scrore=%d", stu.Name, stu.Age, stu.Score)
}
