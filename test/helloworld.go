package main

import (
	"fmt"
	mathClass "golangproject01/myMath"
)

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("hello world1111 ")
	fmt.Println(mathClass.Add(1, 2))
	fmt.Println(mathClass.Sub(4, 2))

	//2， %d 表示整型数字，%s 表示字符串
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)
	//Code=123&endDate=2020-12-31

	//3，声明变量的一般形式是使用 var 关键字：
	//var identifier type;
	var a string = "Runoob"
	fmt.Println(a)

	//4，可以一次声明多个变量：
	//var identifier1, identifier2 type
	var b, c int = 1, 2
	fmt.Println(b, c)

	//5，变量声明
	//第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	//var v_name v_type
	//v_name = value

	//6， 声明一个变量并初始化
	var a1 = "RUNOOB"
	fmt.Println(a1)

	//7， 没有初始化就为零值
	var b1 int
	fmt.Println(b1)

	//8， bool 零值为 false
	var c1 bool
	fmt.Println(c1)

	//9，以下几种类型为 nil：
	/*var a *int
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error // error 是接口
	*/
	var i int
	var f float64
	var b3 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b3, s) //0 0 false ""

	//10，第二种，根据值自行判定变量类型。
	//var v_name = value
	var d = true
	fmt.Println(d) //true

	//11，第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：
	//v_name := value
	//例如：
	/*var intVal int
	intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	*/
	//直接使用下面的语句即可：
	intVal := 1
	fmt.Println(intVal)

	//intVal1 := 1 相等于
	var intVal1 int
	intVal1 = 1
	fmt.Println(intVal1)

	//可以将 var f string = "Runoob" 简写为 f := "Runoob"：
	f1 := "zhudyaang"
	fmt.Println(f1)

	//多变量声明
	//类型相同多个变量, 非全局变量
	/*
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3
	*/
	//var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
	//vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

	//第一种方式
	var vname1, vname2, vname3 string
	vname1, vname2, vname3 = "sss", "kkddd", "dddss"
	fmt.Println(vname1, vname2, vname3)

	//第一种方式
	var vname11, vname22, vname33 = "朱大洋", "里斯", "王武"
	fmt.Println(vname11, vname22, vname33)

	//第一种方式
	vname111, vname222, vname333 := "hust", "wust", "sjtu"
	fmt.Println(vname111, vname222, vname333)

	//这种因式分解关键字的写法一般用于声明全局变量
	/*var (
		vname1 v_type1
		vname2 v_type2
	)
	*/

	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
