package main

import "fmt"

func main() {
	//1,创建channel
	ch := make(chan int, 3)

	//2,给channel存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3,获取channel里面的内容
	a := <-ch
	fmt.Println("", a)

	<-ch //从管道里面取值 //21

	c := <-ch
	fmt.Println("c:", c) //32

	ch <- 56

	//4,打印管道的容量和长度
	fmt.Printf("值:%v 容量:%v 长度:%v\n", ch, cap(ch), len(ch)) //值:0xc00012e000 容量:3 长度:1

	//5，管道的类型（引用数据类型）
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 25

	<-ch1
	<-ch1
	<-ch1

	d := <-ch1
	fmt.Println("d:", d)

	// 管道的阻塞
	/*
			ch6 := make(chan int, 1)
		    ch6 <- 34
		    ch6 <- 64 //fatal error: all goroutines are asleep - deadlock!
	*/

	ch7 := make(chan string, 2)
	ch7 <- "数据1"
	ch7 <- "数据2"

	m1 := <-ch7
	m2 := <-ch7
	m3 := <-ch7
	fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock! 管道没有数据了 你还在取 也会阻塞
}
