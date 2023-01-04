package main

import (
	"fmt"
	"time"
)

// 在主线程中开启一个goroutine,该协程每隔50秒中输出"你好 golang"
// 在主线程中也m每隔一秒输出"你好 golang"，输出10次后 退出程序
// 要求主线程和goroutine同时进行
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好 golang-", i)
		time.Sleep(time.Millisecond * 100)
	}

}
func main() {
	go test() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main()-你好 golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
}

/*可以看到有两个进程：
  主进程main进程 和 协程
  协程会被 main进程 进行控制，就是  协程 比较慢的话，main进程比较快，这时候就会退出程序，不管协程有没有被执行完毕 都会退出

*/
