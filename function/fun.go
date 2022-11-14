package main

import "fmt"

func test(num int) {
	fmt.Println("----test001=", num)
}
func test01(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println(num1, num2, testFunc)

}
func main() {
	//a := test
	//fmt.Printf("a的类型为: %T,test的类型为: %T", a, test) //a的类型为: func(int),test的类型为: func(int)

	//fmt.Println()
	//a(10) //----test001= 10
	//test(10)

	fmt.Println()
	test01(10, 3.45, test)
}
