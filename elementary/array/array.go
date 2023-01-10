package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {

	//数组初始化方式1：
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	var arr2 = [3]string{"java", "php", "golang"}
	fmt.Println(arr2)

	var arr3 = [...]string{"java", "php", "golang", "c++"}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 2, 2: 3}
	fmt.Println(arr4)

	chinese := "mp4/视频-1/4093042-264875cfa9363d5187bc9661c0682bee.mp4"
	str1 := base64.StdEncoding.EncodeToString([]byte(chinese))
	fmt.Println(str1)

	str2, _ := base64.StdEncoding.DecodeString("Y2hpbmEv6Y+J4oKs5rWc77+9LTEvMzYyNjk2NC0yY2NjYjI0YWY4Yjg1YmUxNjczMGMxN2E5ZTJmY2UzMy5tcDQ=")
	fmt.Println("str2:", string(str2))

	str := "bXA0Lz8/LTEvMzYyNjk2NC0yY2NjYjI0YWY4Yjg1YmUxNjczMGMxN2E5ZTJmY2UzMy5tcDQ="
	comma := strings.Index(str, "/")
	fmt.Println(comma)
	result := str[comma+1 : len(str)-1]
	fmt.Println(result)

	ss := strings.Contains(str, "/")
	fmt.Println("ss:", ss)

	test()

}

func test() {
	fmt.Println("-----------------------")
	var arrays = []string{
		"bXA0Lz8/LTEvMzYyNjk2NC0yY2NjYjI0YWY4Yjg1YmUxNjczMGMxN2E5ZTJmY2UzMy5tcDQ=",
		"bXA0Lz8/LTEvNDA5MzA0Mi0yNjQ4NzVjZmE5MzYzZDUxODdiYzk2NjFjMDY4MmJlZS5tcDQ=",
		"bXA0LzQ3OTU4OTAtNGQ3N2U1NDZhNDYzNDQ3OGM5MmU2Y2ViNzk2MGFhM2UubXA0",
		/*	"bXA0LzQ3ODA5ODUtZWEzMjBjODdiNDU2ZjZlYzhhMjhiNTc2YmU5N2VkYzEubXA0",
			"bXA0Lz8/LTIvNDM2MTA1Mi0wOGI0YTlmMWUxOWJkMDFiYWM2MTJjYTMyZDYwMGNkZi5tcDQ=",
			"bXA0Lz8/LTIvNDcxMDY3MS0yZjFjMWM3ODNhMDE5ZjU2Y2U3ODZhMGU5OGEzYzEzZS5tcDQ=",
			"bXA0Lz8/LTEvPz8tMS0xLzI4ODU4MDMtN2E5Yzg0NzY3MGI1YzA5YTJjZDNjMGY3MTYzMTVjOTcubXA0",
			"bXA0Lz8/LTIvNDM1NjY0OS0yYTQ2Y2M5NDY1NDg4YTg1Mjg2NzdjM2Q3MzgzNzdjMy5tcDQ=",*/
	}

	var srcPath []byte
	var nerrRe error

	for _, v := range arrays {
		if find := strings.Contains(v, "/"); find {
			index := strings.Index(v, "/")
			if index != -1 {
				strSrcPath := v[index+1 : len(v)-1]
				srcPath, nerrRe = base64.StdEncoding.DecodeString(strSrcPath)
			}
		} else {
			srcPath, nerrRe = base64.StdEncoding.DecodeString(v)
		}

		fmt.Println("srcPath:", string(srcPath))

		if nerrRe == nil {
			v = string(srcPath)
			//fmt.Println("v:", v)
		}

	}

}
