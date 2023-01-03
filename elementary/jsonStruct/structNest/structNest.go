package main

import (
	"encoding/json"
	"fmt"
)

// 结构体嵌套
type Student struct {
	ID     int
	Gender string
	Name   string
}
type Class struct {
	Title    string
	Students []Student
}

func main() {
	structToJsonNest()
	JsonNestToStruct()

}

func structToJsonNest() {
	c := Class{
		Title:    "001班",
		Students: make([]Student, 0, 200),
	}
	for i := 1; i <= 10; i++ {
		s := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	//fmt.Println("c:", c)

	strByte, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:\n", string(strByte))
}
func JsonNestToStruct() {
	str := `{"Title":"001班","Students":[{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"},{"ID":5,"Gender":"男","Name":"stu_5"},{"ID":6,"Gender":"男","Name":"stu_6"},{"ID":7,"Gender":"男","Name":"stu_7"},{"ID":8,"Gender":"男","Name":"stu_8"},{"ID":9,"Gender":"男","Name":"stu_9"},{"ID":10,"Gender":"男","Name":"stu_10"}]}`

	//方式-：
	/*var c = &Class{}
	err := json.Unmarshal([]byte(str), c)*/

	////方式二：
	var c Class
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", c)
}
