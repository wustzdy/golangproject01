package main

import "fmt"

// 自定义类型
type myInt int

// 自定义类型别名
type myFloat = float64

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	var a myInt = 10
	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T\n", b, b) //12.3 float64

	//使用结构体
	structTest()
}

func structTest() {
	//1，实例化结构体
	var p1 Person //实例化结构体person
	p1.Name = "张三"
	p1.Age = 12
	p1.Sex = "男"
	fmt.Printf("值:%v 类型:%T\n", p1, p1)  //值:{张三 12 男} 类型:main.person
	fmt.Printf("值:%#v 类型:%T\n", p1, p1) //值:main.person{name:"张三", age:12, sex:"男"} 类型:main.person

	//2，可以通过new实例化结构体person
	//注意：在golang中支持对结构体指针直接使用.来访问结构体中的成员，p2.name="张三" 其实在底层是(*p2).name="张三"
	var p2 = new(Person)
	p2.Name = "李四"
	p2.Age = 30
	p2.Sex = "女"

	(*p2).Name = "王哈哈"
	fmt.Printf("值:%#v 类型:%T\n", p2, p2) //值:&main.Person{Name:"李四", Age:30, Sex:"女"} 类型:*main.Person

	//3，第三种实例化结构体person
	var p3 = &Person{
		Name: "社会生活",
		Sex:  "女",
		Age:  34,
	}
	fmt.Printf("值:%#v 类型:%T\n", p3, p3) //值:&main.Person{Name:"社会生活", Age:34, Sex:"女"} 类型:*main.Person

	//4，第四种实例化结构体person
	var p4 = &Person{}
	p4.Name = "李四"
	p4.Age = 30
	p4.Sex = "女"

	fmt.Printf("值:%#v 类型:%T\n", p3, p3) //值:&main.Person{Name:"社会生活", Age:34, Sex:"女"} 类型:*main.Person

	//5，第五种实例化结构体person
	var p5 = Person{
		Name: "社会生活",
		Sex:  "女",
		Age:  34,
	}
	fmt.Printf("值:%#v 类型:%T\n", p5, p5) //值:main.Person{Name:"社会生活", Age:34, Sex:"女"} 类型:main.Person

	//6，第六种实例化结构体person
	var p6 = &Person{
		Name: "社会生活",
	}
	fmt.Printf("值:%#v 类型:%T\n", p6, p6) //值:&main.Person{Name:"社会生活", Age:0, Sex:""} 类型:*main.Person

	//7，第七种实例化结构体person
	var p7 = &Person{
		"找死",
		20,
		"男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p7, p7) //值:&main.Person{Name:"找死", Age:20, Sex:"男"} 类型:*main.Person

	fmt.Println("---验证结构体是否是引用数据类型-----")
	var p8 = Person{
		Name: "社会生活",
		Sex:  "女",
		Age:  34,
	}
	p9 := p8
	fmt.Printf("p8 %#v\n", p8) //p8 main.Person{Name:"社会生活", Age:34, Sex:"女"}
	fmt.Printf("p9 %#v\n", p9) //p9 main.Person{Name:"社会生活", Age:34, Sex:"女"}

	p9.Name = "武汉"
	fmt.Printf("p8 %#v\n", p8) //p8 main.Person{Name:"社会生活", Age:34, Sex:"女"}
	fmt.Printf("p9 %#v\n", p9) //p9 main.Person{Name:"武汉", Age:34, Sex:"女"}
	//说明结构体是一个值类型
}
