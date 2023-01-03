package main

import "fmt"

// 类型断言
func main() {
	alert()
	test()

}

func test() {
	MyPrint("你好golang")
	MyPrint1("你好golang")
}

func alert() {
	var a interface{}
	a = "你好golang"
	v, ok := a.(string) //

	if ok {
		fmt.Println("a 就是一个string类型，值是:\n", v) //a 就是一个string类型，值是: 你好golang
	} else {
		fmt.Println("断言失败\n")
	}
}

// 定义一个方法，可以传入任意数据类型，然后根据不同类型实现不同的功能
func MyPrint(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("a 就是一个string类型，值是:", v)
	} else if _, ok := x.(int); ok {
		fmt.Println("a 就是一个int类型，值是:", v)
	} else if _, ok := x.(bool); ok {
		fmt.Println("a 就是一个bool类型，值是:", v)
	}
}

// x.(type)判断一个变量的类型，只能用于switch语句中
func MyPrint1(x interface{}) {
	fmt.Println("---------------")
	switch x.(type) {
	case int:
		fmt.Println("a 就是一个int类型")
	case string:
		fmt.Println("a 就是一个string类型")
	case bool:
		fmt.Println("a 就是一个bool类型")
	default:
		fmt.Println("传入错误")
	}
}
