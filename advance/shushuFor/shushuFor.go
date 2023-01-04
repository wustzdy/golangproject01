package main

import "fmt"

// 统计1-100哪些是素数 for循环
func main() {
	for num := 2; num <= 100; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num, "是素数!")
		}
	}
}
