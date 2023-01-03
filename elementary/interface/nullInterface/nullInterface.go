package main

import "fmt"

type A interface{} //空接口，表示没有约束，任意类型的接口都可以实现空接口

func main() {
	nullInterfaceTest()
	nullFunc()
	nullShow()

}

func nullShow() {
	show(20)
	show("你好golang")
	slice := []int{1, 2, 3, 4}
	show(slice) //传入一个切片类型
}

func nullInterfaceTest() {
	var a A
	var str = "你好golang"
	a = str                           //让字符串实现A这个接口
	fmt.Printf("值:%v  类型:%T\n", a, a) //值:你好golang  类型:string

	var num = 20
	a = num                           //表示让int类型实现A这个接口
	fmt.Printf("值:%v  类型:%T\n", a, a) //值:20  类型:int

	var flag = true
	a = flag
	fmt.Printf("值:%v  类型:%T\n", a, a) //值:true  类型:bool

	fmt.Println("-----------------")
}

func nullFunc() {
	//golang中空接口也可以直接当做类型来使用，可以表示任意类型
	var a interface{}
	a = 20
	fmt.Printf("值:%v  类型:%T\n", a, a) ///值:20  类型:int
	a = "你好golang"
	fmt.Printf("值:%v  类型:%T\n", a, a) //值:你好golang  类型:string
	a = true
	fmt.Printf("值:%v  类型:%T\n", a, a) //值:true  类型:bool

	fmt.Println("-----------------")
}

// //golang中空接口也可以直接当做类型来使用，可以表示任意类型
func show(a interface{}) {
	fmt.Printf("值:%v  类型:%T\n", a, a)
}
