package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//我们想在切片里面存放一系列的用户信息，这时候我们可以定义一个元素为map类型的切片
	var userInfo = make([]map[string]string, 3, 3)
	fmt.Println(userInfo[0])        //map[]  map的默认值是nil
	fmt.Println(userInfo[0] == nil) //true

	if userInfo[0] == nil {
		userInfo[0] = make(map[string]string)
		userInfo[0]["username"] = "张三"
		userInfo[0]["age"] = "12"
		userInfo[0]["height"] = "189cm"
	}

	if userInfo[1] == nil {
		userInfo[1] = make(map[string]string)
		userInfo[1]["username"] = "李四"
		userInfo[1]["age"] = "22"
		userInfo[1]["height"] = "172cm"
	}

	if userInfo[2] == nil {
		userInfo[2] = make(map[string]string)
		userInfo[2]["username"] = "王武"
		userInfo[2]["age"] = "22"
		userInfo[2]["height"] = "123cm"
	}
	fmt.Println("userInfo:", userInfo) // [map[age:12 height:189cm username:张三] map[age:22 height:172cm username:李四] map[age:22 height:123cm username:王武]]

	for k, v := range userInfo {
		fmt.Println(k, v)
	}
	/*
		0 map[age:12 height:189cm username:张三]
		1 map[age:22 height:172cm username:李四]
		2 map[age:22 height:123cm username:王武]
	*/
	fmt.Println("------------------循环-----")
	for _, v := range userInfo {
		for key, value := range v {
			fmt.Printf("key:%v value:%v\n", key, value)
		}
	}
	/*
		key:username value:张三
		key:age value:12
		key:height value:189cm

		key:username value:李四
		key:age value:22
		key:height value:172cm

		key:username value:王武
		key:age value:22
		key:height value:123cm
	*/

	var userInfo1 = make(map[string]string)
	userInfo1["username"] = "张三"
	userInfo1["hobby"] = "吃饭"

	var userInfo2 = make(map[string][]string)
	userInfo2["hobby"] = []string{
		"吃饭",
		"睡觉",
		"打豆豆",
	}
	userInfo2["work"] = []string{
		"java",
		"golang",
		"C",
	}
	fmt.Println(userInfo2)

	for _, v := range userInfo2 {
		for key, value := range v {
			fmt.Println("", key, value)
		}
	}

	//map也是引用数据类型
	var userInfoM = make(map[string]string)
	userInfoM["username"] = "张三"
	userInfoM["age"] = "20"

	userInfoN := userInfoM
	fmt.Println("userInfoM:", userInfoM) //userInfoM: map[age:20 username:张三]
	fmt.Println("userInfoN", userInfoN)  //userInfoN map[age:20 username:张三]

	userInfoN["username"] = "李四"
	fmt.Println("userInfoM:", userInfoM) //userInfoM: map[age:20 username:李四]
	fmt.Println("userInfoN", userInfoN)  //userInfoN map[age:20 username:李四]
	//说明改变副本会影响原来的值，说明map是引用数据类型

	//map排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	//按照key的升序输出
	//1，把map的key放在切片里面
	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	fmt.Println("keySlice:", keySlice) //keySlice: [10 4 1 8]
	//2，对key进行升序排序
	sort.Ints(keySlice)
	fmt.Println("升序后keySlice:", keySlice) //升序后keySlice: [1 4 8 10]
	//2，循环遍历key
	for _, v := range keySlice {
		fmt.Printf("key=%v value=%v\n", v, map1[v])
	}
	/*
		key=1 value=13
		key=4 value=56
		key=8 value=90
		key=10 value=100
	*/

	var str = "how do you do"
	strSlice := strings.Split(str, " ")
	fmt.Println("strSlice:", strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap) //map[do:2 how:1 you:1]
}
