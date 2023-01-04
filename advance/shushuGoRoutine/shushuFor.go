package main

import (
	"fmt"
	"time"
)

// 统计1-100哪些是素数 for循环
func main() {
	start := time.Now().Unix()
	for num := 2; num <= 120000; num++ {
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
	end := time.Now().Unix()
	fmt.Println("一共运行的时间：", end-start) //总耗时4秒中

}
