package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

func main() {
	var p Person
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "写代码"
	p.Hobby[1] = "打架"
	p.Hobby[2] = "杀人"

	p.map1 = make(map[string]string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "12345566"

	fmt.Printf("%#v\n", p) //main.Person{Name:"张三", Age:20, Hobby:[]string{"写代码", "打架", "杀人"}, map1:map[string]string{"address":"北京", "phone":"12345566"}}

	fmt.Printf("%v", p.Hobby) //[写代码 打架 杀人]
}
