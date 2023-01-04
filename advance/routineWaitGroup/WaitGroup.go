package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好 golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器减1
}
func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好 golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器减1
}
func main() {
	wg.Add(1)  //协程计数器加1
	go test1() //表示开启一个协程

	wg.Add(1)  //协程计数器加1
	go test2() //表示开启一个协程

	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程结束")
}
