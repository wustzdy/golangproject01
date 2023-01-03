package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	structToStrTag()

}

func structToStrTag() {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s001",
	}
	fmt.Printf("%#v\n", s1) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s001"}

	jsonByte, _ := json.Marshal(s1)
	fmt.Println("result:", string(jsonByte)) //result: {"id":12,"gender":"男","name":"李四","sno":"s001"}
}
