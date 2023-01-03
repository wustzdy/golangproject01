package main

import (
	"fmt"
	"sort"
)

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

// 升序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

// 降序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

func mapSort(map1 map[string]string) string {
	//1，把map对象的key放在一个map的切片里面
	var sliceKey []string
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}
	fmt.Println(sliceKey)

	//2，对key进行升序排序
	sort.Strings(sliceKey)
	fmt.Println("排序后的sliceKey:", sliceKey)

	//3,
	var str string
	for _, value := range sliceKey {
		str += fmt.Sprintf("%v=>%v ", value, map1[value])
	}
	return str

}

// 递归打印整数
func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

// 递归实现1-100的和
func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}
}

// 递归实现5的阶乘
func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}

// 闭包:函数里面嵌套一个函数 最后返回里面的函数
// 1，可以让一个变量常驻内存
// 2，可以让一个变量不污染全局
func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
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

	//升序
	var sliceA = []int{12, 34, 37, 35, 556, 36, 2}
	arr := sortIntAsc(sliceA)
	fmt.Println("arr:", arr) //arr: [2 12 34 35 36 37 556]

	//降序
	var sliceB = []int{12, 34, 37, 35, 556, 36, 2}
	arrB := sortIntDesc(sliceB)
	fmt.Println("arrB:", arrB) //arrB: [556 37 36 35 34 12 2]

	//对map的key进行升序排序
	map1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"heigh":    "180cm",
	}

	str := mapSort(map1)
	fmt.Println("str:", str) //str: age=>20 heigh=>180cm sex=>男 username=>zhangsan

	fn1(10)
	fmt.Println(fn2(100)) //5050
	fmt.Println(fn3(5))   //120

	var a = adder()  //表示执行方法
	fmt.Println(a()) //11
	fmt.Println(a()) //11
	fmt.Println(a()) //11

	var fn2 = adder2()   //表示执行方法
	fmt.Println(fn2(10)) //20
	fmt.Println(fn2(10)) //30
	fmt.Println(fn2(10)) //40
}
