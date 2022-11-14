package main

import (
	"fmt"
	"math"
)

func number() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

	getNumber := number()
	fmt.Println(getNumber())
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "zhangsan"
	return a, b, c
}
