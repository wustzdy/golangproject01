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
func timeFormat() {
	//时间戳转为时分秒
	unixTime := 1672660651
	timeObj := time.Unix(int64(unixTime), 0)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println("str:", str) //str: 2023-01-02 19:57:31
}
