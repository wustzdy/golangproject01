package main

import "fmt"

func main() {
	var a = 10
	fmt.Printf("a的值:%v a的类型:%T a的地址:%p\n", a, a, &a) //a的值:10 a的类型:int a的地址:0xc00001a098

	//指针也是一个变量，但是他是一个特殊的变量，它存储的数据不是一个普通的值，而是另外一个变量的内存地址
	var b = 10
	var p = &b                           //p指针变量 p的类型 *int
	fmt.Printf("p的值:%v p的类型:%T\n", p, p) //p的值:0xc000124018 p的类型:*int

	pointerTest()
	pointerTest1()

	fmt.Println("-------------------") //30
	var aa = 5
	fn1(aa)
	fmt.Println(aa) //5

	fn2(&aa)
	fmt.Println(aa) //40

	mapTest()
}

func pointerTest() {
	//golang 里面的变量都有一个内存地址
	var a = 10
	var b = &a                                       //b指针变量 b的类型 *int
	fmt.Printf("a的值:%v a的类型:%T a的地址:%p\n", a, a, &a) //a的值:10 a的类型:int a的地址:0xc000124030
	fmt.Printf("b的值:%v b的类型:%T b的地址:%p\n", b, b, &b) //b的值:0xc000124030 b的类型:*int b的地址:0xc00011c020
	//可以看出p的值就是a的内存地址，但是b也有自己的内存地址
}

func pointerTest1() {
	a := 10
	p := &a //p指针变量 类型* int
	//*p 表示取出p这个变量对应的内存地址的值
	fmt.Println(p)  //0xc00001a0c8  //a的地址
	fmt.Println(*p) //10 //*p表示打印出a的值

	*p = 30        //改变p这个变量对应的内存地址的值
	fmt.Println(a) //30
}

func fn1(x int) {
	x = 10
}
func fn2(x *int) {
	*x = 40
}

func mapTest() {

	//以下写法不对
	/*var userInfo = map[string]string
	userInfo["username"] = "张三"
	userInfo["sex"] = "男"
	fmt.Println(userInfo)*/
	//因为map是引用数据类型，使用之前要先进行分配空间就是用make

	var userInfo = make(map[string]string)
	userInfo["username"] = "张三"
	userInfo["sex"] = "男"
	fmt.Println(userInfo) //map[sex:男 username:张三]

	/*var a *int
	*a = 100
	fmt.Println(*a)*/ //也会报错，因为指针也是一个引用数据类型

	//我们可以用new 来分配空间 使用new函数得到的是一个指针类型，并且改指针对应的值为该类型d的零值

	var a = new(int)                                     //表示a是一个指针变量， 类型是*int的指针类型 值是0
	fmt.Printf("a的值:%v a的类型:%T 指针变量对应的值:%v\n", a, a, *a) //a的值:0xc000122040 a的类型:*int 指针变量对应的值:0

	/*
		错误的写法
		var a *int
		*a = 100
		fmt.Println(*a)
	*/
	var aa *int
	aa = new(int)
	*aa = 100
	fmt.Println(*aa) //100
}
