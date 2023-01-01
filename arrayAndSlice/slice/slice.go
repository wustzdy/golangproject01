package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {

	var arr1 = []int{1, 2, 3, 4}
	fmt.Printf("%v - %T - 长度:%v\n", arr1, arr1, len(arr1))

	var arr2 = []int{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Printf("%v - %T - 长度:%v\n", arr2, arr2, len(arr2))

	var arr3 []int
	fmt.Println(arr3 == nil) //说明切片的默认值是nil

	var strSlice = [...]string{"java", "php", "golang", "c++"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	for k, v := range strSlice {
		fmt.Println(k, v)
	}

	//基于数组切片
	//a := [5]int{}

	m := map[string]string{"userID": "111", "eee": "22"}
	fmt.Println("m:", m)
	//json := `{"indexCreating":false,"list":[{"id":1,"createdAt":1671441014,"fileType":0,"url":"https://10.198.30.201/oss/spe_data/test/datasets/v1hy/files/05520eea-58ab-4328-967c-fb6abd1d455a.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=spe%2F20221226%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20221226T033423Z\u0026X-Amz-Expires=3600\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=691d7f9cc5f67a801fd9e3fa836205a6448c2e7439f627371874c7bf8bfd847d","filename":"cat.10007.jpg","size":10328,"hasResult":true,"result":"{\"height\":300,\"rotate\":0,\"step_1\":{\"dataSourceStep\":0,\"toolName\":\"rectTool\",\"result\":[{\"x\":39.22457504272461,\"y\":0,\"width\":124.2020263671875,\"height\":293.3439025878906,\"order\":2,\"valid\":true,\"id\":\"0da2d43d-78e0-474f-a56e-babd0f3e12ee\",\"sourceID\":\"0\",\"attribute\":\"cat\",\"textAttribute\":\"\"},{\"x\":71.17144012451172,\"y\":101.07936096191406,\"width\":103.14998626708984,\"height\":183.77952575683594,\"order\":3,\"valid\":true,\"id\":\"08c01e7e-4dc3-467e-bac8-2801f50ede66\",\"sourceID\":\"0\",\"attribute\":\"cat\",\"textAttribute\":\"\"},{\"x\":157.60549926757812,\"y\":119.9735336303711,\"width\":31.373245239257812,\"height\":126.6599349975586,\"order\":11,\"valid\":true,\"id\":\"67f5b9c2-0769-4e34-85ab-494390d18f7e\",\"sourceID\":\"0\",\"attribute\":\"carrot\",\"textAttribute\":\"\"}]},\"totalFrames\":0,\"valid\":true,\"width\":219}","meta":{"imageHeight":300,"imageWidth":219,"md5":"f44da808e68eb337eeeb9fc905a972af"},"confidenceMap":"{\"08c01e7e-4dc3-467e-bac8-2801f50ede66\":0.3813336,\"0da2d43d-78e0-474f-a56e-babd0f3e12ee\":0.9526289,\"67f5b9c2-0769-4e34-85ab-494390d18f7e\":0.3206764}"}],"progress":0,"total":1986}`
	//json := `{"indexCreating":false,"list":[],"progress":0,"total":0}`
	json := `{
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"age": 44, "first": "Dale", "last": "Murphy"},
    {"age": 68, "first": "Roger", "last": "Craig"},
    {"age": 47, "first": "Jane", "last": "Murphy"}
  ],
  "name": {"first": "Tom", "last": "Anderson"}
}`

	json = `{
    "indexCreating":false,
    "list":[
        {
            "id":1,
            "createdAt":1671441014,
            "fileType":0,
            "url":"https://10.198.30.201/oss/spe_data/test/datasets/v1hy/files/05520eea-58ab-4328-967c-fb6abd1d455a.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=spe%2F20221226%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20221226T033423Z\u0026X-Amz-Expires=3600\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=691d7f9cc5f67a801fd9e3fa836205a6448c2e7439f627371874c7bf8bfd847d",
            "filename":"cat.10007.jpg",
            "size":10328,
            "hasResult":true,
            "result":"{\"height\":300,\"rotate\":0,\"step_1\":{\"dataSourceStep\":0,\"toolName\":\"rectTool\",\"result\":[{\"x\":39.22457504272461,\"y\":0,\"width\":124.2020263671875,\"height\":293.3439025878906,\"order\":2,\"valid\":true,\"id\":\"0da2d43d-78e0-474f-a56e-babd0f3e12ee\",\"sourceID\":\"0\",\"attribute\":\"cat\",\"textAttribute\":\"\"},{\"x\":71.17144012451172,\"y\":101.07936096191406,\"width\":103.14998626708984,\"height\":183.77952575683594,\"order\":3,\"valid\":true,\"id\":\"08c01e7e-4dc3-467e-bac8-2801f50ede66\",\"sourceID\":\"0\",\"attribute\":\"cat\",\"textAttribute\":\"\"},{\"x\":157.60549926757812,\"y\":119.9735336303711,\"width\":31.373245239257812,\"height\":126.6599349975586,\"order\":11,\"valid\":true,\"id\":\"67f5b9c2-0769-4e34-85ab-494390d18f7e\",\"sourceID\":\"0\",\"attribute\":\"carrot\",\"textAttribute\":\"\"}]},\"totalFrames\":0,\"valid\":true,\"width\":219}",
            "meta":{
                "imageHeight":300,
                "imageWidth":219,
                "md5":"f44da808e68eb337eeeb9fc905a972af"
            },
            "confidenceMap":"{\"08c01e7e-4dc3-467e-bac8-2801f50ede66\":0.3813336,\"0da2d43d-78e0-474f-a56e-babd0f3e12ee\":0.9526289,\"67f5b9c2-0769-4e34-85ab-494390d18f7e\":0.3206764}"
        }
    ],
    "progress":0,
    "total":1986
}`
	value := gjson.Get(json, "list.#.result")
	str := strings.Replace(value.String(), "\n", "", -1)
	println("ssss:" + str)

	/*result := gjson.Get(json, "list")
	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true // keep iterating
	})*/

	a := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	b := a[1:]
	fmt.Println("b:", b) //[上海 广州 深圳 成都 重庆]

	//3,关于切片的长度和容量
	//长度：切片的长度就是它所包含元素的个数
	//容量：切片的容量是从它的d第一个元素开始数，到其底层数组元素末尾的个数
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) //长度6 容量6

	s1 := s[2:]                                 //5, 7, 11, 13
	fmt.Printf("长度%d 容量%d\n", len(s1), cap(s1)) //长度4 容量4

	s2 := s[1:3]                                //3,5
	fmt.Printf("长度%d 容量%d\n", len(s2), cap(s2)) //长度2 容量5

	s3 := s[:3]                                 //2, 3, 5
	fmt.Printf("长度%d 容量%d\n", len(s3), cap(s3)) //长度3 容量6

	//4，使用make来声明一个切片
	var sliceA = make([]int, 4, 8)
	fmt.Println(sliceA)                                 //[0 0 0 0]
	fmt.Printf("长度%d 容量%d\n", len(sliceA), cap(sliceA)) //长度4 容量8

	var sliceB = make([]int, 4, 8)
	sliceB[0] = 10
	sliceB[1] = 12
	sliceB[2] = 40
	sliceB[3] = 30
	fmt.Println(sliceB) //[10 12 40 30]

	slicec := []string{"php", "java", "go"}
	slicec[2] = "golang"
	fmt.Println(slicec) //[php java golang]

	//golang 中没法通过下标的方式给切片进行扩容，需要用到append()
	/*var sliced []int
	fmt.Printf("长度%d 容量%d\n", len(sliced), cap(sliced)) //长度0 容量0
	sliced[0] = 1
	fmt.Println(sliced)*/

	//1，golang 要给切片扩容的话要用append方法
	var sliceE []int
	fmt.Printf("长度%d 容量%d\n", len(sliceE), cap(sliceE)) //长度0 容量0
	sliceE = append(sliceE, 12)
	sliceE = append(sliceE, 24)
	fmt.Printf("%v  长度%d 容量%d\n", sliceE, len(sliceE), cap(sliceE)) //[12 24]  长度2 容量2

	//2，可以一次性追加多个元素
	var sliceF []int
	fmt.Printf("长度%d 容量%d\n", len(sliceF), cap(sliceF)) //长度0 容量0
	sliceF = append(sliceF, 12, 23, 35, 465)
	fmt.Printf("%v 长度%d 容量%d\n", sliceF, len(sliceF), cap(sliceF)) //[12 23 35 465] 长度4 容量4

	//3，append还可以合并切片
	sliceG := []string{"php", "java"}
	sliceH := []string{"nodeJs", "golang"}

	sliceG = append(sliceG, sliceH...)                             //把sliceH合并到sliceG上
	fmt.Printf("%v 长度%d 容量%d\n", sliceG, len(sliceG), cap(sliceG)) //[php java nodeJs golang] 长度4 容量4

	//4，切片的扩容策略
	var sliceI []int
	for i := 1; i <= 10; i++ {
		sliceI = append(sliceI, i)
		fmt.Printf("%v 长度%d 容量%d\n", sliceI, len(sliceI), cap(sliceI))
	}

	//5，切片就是引用数据类型
	sliceJ := []int{1, 2, 3, 45}
	sliceK := sliceJ
	sliceK[0] = 11
	fmt.Println(sliceJ) //[11 2 3 45]
	fmt.Println(sliceK) //[11 2 3 45]

	//6，使用copy来复制切片，sliceN改变了，但是不会影响sliceM的切片
	sliceM := []int{1, 2, 3, 45}
	sliceN := make([]int, 4, 4)
	copy(sliceN, sliceM)
	fmt.Println(sliceN) //[1 2 3 45]
	fmt.Println(sliceM) //[1 2 3 45]

	sliceN[0] = 111111
	fmt.Println(sliceN) //[111111 2 3 45]
	fmt.Println(sliceM) //[1 2 3 45]

	//7，删除切片元素的方法
	a1 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//要删除索引为2的元素
	a1 = append(a1[:2], a1[3:]...)
	fmt.Printf("%v 长度%d 容量%d\n", a1, len(a1), cap(a1)) //[30 31 33 34 35 36 37] 长度7 容量8

	//8-改变字符串中某一个字符
	s11 := "big"
	byteStr := []byte(s11)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr)) //pig

	s22 := "你好golang"
	runeStr := []rune(s22)
	runeStr[0] = '大'
	fmt.Println(string(runeStr)) //大好golang
}
