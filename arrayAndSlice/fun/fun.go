package main

import "fmt"

// golang中可变参数
func sumFu1(x ...int) int {
	fmt.Printf("%v %T\n", x, x) //[12 34 45 46] []int

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 固定参数和可变参数
func sumFu2(x int, y ...int) int {
	fmt.Println(x, y) //100 [1 2 3 4]

	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func fun3(x, y int) int {
	return x + y
}

// 多个返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc1(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub) //0 0
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub) //12 8
	return
}

func calc2(x, y int) (sum, sub int) {
	fmt.Println(sum, sub) //0 0
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub) //12 8
	return
}
func main() {
	sum1 := sumFu1(12, 34, 45, 46)
	fmt.Println("sum1:", sum1)

	sum2 := sumFu2(100, 1, 2, 3, 4)
	fmt.Println("sum2:", sum2) //sum2: 110

	sum3 := fun3(10, 2)
	fmt.Println("sum3:", sum3) //sum3: 12

	sum, sub := calc(10, 2)
	fmt.Printf("sum:%d sub:%d\n", sum, sub) //sum:12 sub:8

	sum11, sub11 := calc1(10, 2)
	fmt.Printf("sum11:%d sub11:%d\n", sum11, sub11) //sum11:12 sub11:8

}
