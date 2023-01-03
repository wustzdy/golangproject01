package main

import "fmt"

func main() {
	//1,创建map类型的数据
	var userInfo = make(map[string]string)
	userInfo["username"] = "张三"
	userInfo["age"] = "20"
	userInfo["sex"] = "男"
	fmt.Println(userInfo)             //map[age:20 sex:男 username:张三]
	fmt.Println(userInfo["username"]) //张三

	//2.map也可以填充数据
	var userInfo1 = map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
	}
	fmt.Println("userInfo1:", userInfo1)

	//3.map也可以填充数据
	userInfo2 := map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
	}
	fmt.Println("userInfo2:", userInfo2)

	//4.map循环-for range
	var userInfo4 = map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
	}
	//fmt.Println(userInfo4["username"]) //里斯

	for k, v := range userInfo4 {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
	/*key:age value:22
	  key:sex value:女
	  key:username value:里斯*/

	//4.map类型数据的crud
	//4.1创建 修改map类型的数据
	var userInfo5 = make(map[string]string)
	userInfo5["username"] = "张三"
	userInfo5["age"] = "20"
	fmt.Println("userInfo5:", userInfo5)

	var userInfo6 = map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
	}
	userInfo6["username"] = "李四"
	fmt.Println("userInfo6:", userInfo6) // map[age:22 sex:女 username:李四]

	//4.2获取查找map类型的数据
	var userInfo7 = map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
	}
	userInfo7["username"] = "李四"
	fmt.Println("userInfo6:", userInfo7["username"]) //获取  userInfo6: 李四

	//4.3查找
	v, ok := userInfo7["age"]
	if ok {
		fmt.Println(v) //22 true
	} else {
		fmt.Println("查不到")
	}

	v1, ok1 := userInfo7["age11"]
	fmt.Println(v1, ok1) //空和false

	//5,删除map类型数据里面的key以及对应的值
	var userInfo8 = map[string]string{
		"username": "里斯",
		"age":      "22",
		"sex":      "女",
		"height":   "180",
	}
	fmt.Println("userInfo8:", userInfo8) //userInfo8: map[age:22 height:180 sex:女 username:里斯]

	delete(userInfo8, "sex")             //删除sex属性
	fmt.Println("userInfo8:", userInfo8) //userInfo8: map[age:22 height:180 username:里斯]

}
