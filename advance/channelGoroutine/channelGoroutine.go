package main

import (
	"fmt"
	"sync"
	"time"
)

// 演示goroutine和channel 协作进行任务
var wg sync.WaitGroup

// 写数据
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据 %v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取】数据 %v成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}
func main() {
	var ch = make(chan int, 10)
	wg.Add(1)
	go fn1(ch)

	wg.Add(1)
	go fn2(ch)

	wg.Wait()
	fmt.Println("退出---")
	/*
		【写入】数据 1成功
		【读取】数据 1成功
		【写入】数据 2成功
		【读取】数据 2成功
		【写入】数据 3成功
		【读取】数据 3成功
		【写入】数据 4成功
		【读取】数据 4成功
		【写入】数据 5成功
		【读取】数据 5成功
		【写入】数据 6成功
		【读取】数据 6成功
		【写入】数据 7成功
		【读取】数据 7成功
		【写入】数据 8成功
		【读取】数据 8成功
		【写入】数据 9成功
		【读取】数据 9成功
		【写入】数据 10成功
		【读取】数据 10成功
		退出---


	*/
}
