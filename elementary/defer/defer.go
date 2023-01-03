package main

import "fmt"

func f1() {
	fmt.Println("开始----")
	defer func() {
		fmt.Println("aaaaa")
	}()
	fmt.Println("结束----")
}

// 匿名返回值
func f2() int {
	var a int //0
	defer func() {
		a++
	}()
	return a //
}

// 命名返回值
func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}
func main() {
	fmt.Println("开始")
	defer fmt.Println("1") //被defer定义的语句 会延迟处理
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("结束")
	/*
		开始
		结束
		3
		2
		1
	*/

	f1()
	fmt.Println("f2():", f2()) //f2(): 0
	fmt.Println("f3():", f3()) //f3(): 1
}
