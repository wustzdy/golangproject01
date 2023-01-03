package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	structToJson()
	JsonToStruct()

}

// 结构体转为字符串
func structToJson() {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s001",
	}
	fmt.Printf("%#v\n", s1) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s001"}

	jsonByte, _ := json.Marshal(s1)
	fmt.Println("result:", string(jsonByte)) //result: {"ID":12,"Gender":"男","Name":"李四","Sno":"s001"}
}
func JsonToStruct() {
	//把字符串转为struct结构体
	//var str = "{\"ID\":12,\"Gender\":\"男\",\"Name\":\"李四\",\"Sno\":\"s001\"}"
	var str = `{"ID":12,"Gender":"男","Name":"李四","Sno":"s001"}`
	var s2 Student
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s2) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s001"}
}
