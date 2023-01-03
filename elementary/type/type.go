package main

import "fmt"

// 自定义类型
type myInt int
type calc func(int, int) int //定义一个calc类型

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// 函数
// 函数作为另外一个函数参数
func add1(x, y int) int {
	return x + y
}
func sub1(x, y int) int {
	return x - y
}

// 自定义一个方法类型
type calcType func(int, int) int

/*func calc1(x, y int, cb func(int, int) int) int {
}*/

func calc1(x, y int, cb calcType) int {
	return cb(x, y)
}

// 函数作为返回值-start
func add2(x, y int) int {
	return x + y
}
func sub2(x, y int) int {
	return x - y
}

type calcType1 func(int, int) int

func do(o string) calcType1 {
	switch o {
	case "+":
		return add2
	case "-":
		return sub2
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

// 函数作为返回值-end

func main() {
	var c calc
	c = sub
	fmt.Printf("c的类型: %T\n", c) //c的类型: main.calc

	fmt.Println(c(10, 5)) //5

	f := sub
	fmt.Printf("f的类型: %T\n", f) //f的类型: func(int, int) int
	fmt.Println(c(15, 5))       //10

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型: %T\n", a) //a的类型: int
	fmt.Printf("b的类型: %T\n", b) //b的类型: main.myInt

	//做加法运算
	fmt.Println(a + int(b)) //30

	//自定义函数做加法
	sum11 := calc1(10, 5, add1)
	fmt.Println("sum11:", sum11) //sum11: 15

	//自定义函数做减法
	sub11 := calc1(10, 5, sub1)
	fmt.Println("sub11:", sub11) //sub11: 5

	//匿名函数
	j := calc1(10, 5, func(x, y int) int {
		return x * y
	})
	fmt.Println("相乘:", j) //相乘: 50

	//调用
	var a1 = do("+")
	fmt.Println(a1(12, 4)) //16

	var a2 = do("*")
	fmt.Println(a2(3, 4)) //12

	//匿名函数 匿名自执行函数
	func() {
		fmt.Println("test匿名函数")
	}() //test匿名函数

	//匿名函数
	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 4)) //8

	//匿名函数 接受参数
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20) //30
}
