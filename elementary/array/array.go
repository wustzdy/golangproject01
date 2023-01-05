package main

import (
	"encoding/base64"
	"fmt"
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
	str2, _ := base64.StdEncoding.DecodeString(str1)
	fmt.Println(string(str2))

}
