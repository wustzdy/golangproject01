package main

import "fmt"

// 循环遍历管道数据
// 使用for range遍历管道，当管道被关闭的时候 就会退出for range,如果没有关闭管道 就会报错 fatal error:
func main() {
	//noCloseChannel()
	closeChannel()
	forChannel()
}

// 通过for循环来遍历管道，可以不用close
func forChannel() {
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
}

// 关闭管道
func closeChannel() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭管道

	//循环遍历管道的值
	//注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}
}

// 不关闭管道 就会报错
func noCloseChannel() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	//循环遍历管道的值
	//注意：管道没有key
	for v := range ch1 {
		fmt.Println(v) //fatal error: all goroutines are asleep - deadlock!
	}
}
