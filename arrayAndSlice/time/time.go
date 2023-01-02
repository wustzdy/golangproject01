package main

import (
	"fmt"
	"time"
)

func main() {
	//打印当前日期
	printNowTime()
	//时间相关格式
	formatTime()
	//时间戳转为年月日时分秒
	timeFormat()
	//把年月日转成时间戳
	formatStr()
	//时间操作函数
	operationTime()
	//golang定时器
	//timer()
	// 执行5次就停止
	timer1()
	//timeSleep
	timeSleep()

}

func printNowTime() {
	timeObject := time.Now()
	fmt.Println("timeObject:", timeObject) //2023-01-02 19:29:12.691242 +0800 CST m=+0.000099214

	//2023-01-02 19:29:12
	year := timeObject.Year()
	month := timeObject.Month()
	day := timeObject.Day()
	hour := timeObject.Hour()
	minute := timeObject.Minute()
	second := timeObject.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second) //2023-1-2 19:32:30

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2023-01-02 19:33:45
}
func formatTime() {
	timeObject := time.Now()
	//2020-04-05
	var str = timeObject.Format("2006-01-02 03:04:05") //golang中必须要要用这个时间格式-12小时制
	fmt.Println("格式化日期str:", str)                      //格式化日期str: 2023-01-02 07:42:35

	var str1 = timeObject.Format("2006-01-02 15:04:05") //golang中必须要要用这个时间格式-24小时制
	fmt.Println("格式化日期str1:", str1)                     //格式化日期str1: 2023-01-02 19:54:33

	unixTime := timeObject.Unix()           //当前时间戳
	fmt.Println("当前时间戳unixTime:", unixTime) //当前时间戳unixTime: 1672660651

	unixNaTime := timeObject.UnixNano()
	fmt.Println("当前纳秒时间戳unixNaTime:", unixNaTime) //当前纳秒时间戳unixNaTime: 1672660735589579000

}

// 时间戳转为年月日时分秒
func timeFormat() {
	//时间戳转为时分秒
	unixTime := 1672660651
	timeObj := time.Unix(int64(unixTime), 0)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println("str:", str) //str: 2023-01-02 19:57:31
}

// 把年月日转成时间戳
func formatStr() {
	t1 := "2023-01-02 19:57:31"
	timeTemplate := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate, t1, time.Local)
	fmt.Println("当前时间转为时间戳：", stamp.Unix()) //当前时间转为时间戳： 1672660651
}

func operationTime() {
	var timeObj = time.Now()
	fmt.Println("timeObj:", timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println("timeObj:", timeObj)
}

// 每隔一秒钟执行一次操作
func timer() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println("t:", t)
	}
}

// 执行5次就停止
func timer1() {
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		n--
		fmt.Println("t:", t)
		if n == 0 {
			ticker.Stop() //终止这个定时器
			break
		}
	}
}
func timeSleep() {
	fmt.Println("aaa")
	time.Sleep(time.Second)
	fmt.Println("aaa2")
	time.Sleep(time.Second)
	fmt.Println("aaa3")
	time.Sleep(time.Second * 5)
	fmt.Println("aaa4")

	//死循环 每隔一秒钟执行定时任务
	for {
		time.Sleep(time.Second)
		fmt.Println("我在执行定时任务")
	}
}
