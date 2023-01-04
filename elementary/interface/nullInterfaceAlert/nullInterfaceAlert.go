package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// golang中空接口和类型断言使用细节
func main() {
	var userInfo = make(map[string]interface{})
	userInfo["username"] = "张三"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"杀人", "打球"}
	//fmt.Println("userInfo:", userInfo["hobby"][1]) //invalid operation: cannot index userInfo["hobby"] (map index expression of type interface{})
	var address = Address{
		Name:  "李四",
		Phone: 155555555,
	}

	userInfo["address"] = address
	fmt.Println(userInfo["address"]) //{李四 155555555}

	/*
		var name = userInfo["address"].Name
		fmt.Printf(name) //userInfo["address"].Name undefined (type interface{} has no field or method Name)
	*/
	hobby2, _ := userInfo["hobby"].([]string)
	fmt.Println("", hobby2[1]) //打球

	address2, _ := userInfo["address"].(Address)
	fmt.Println(address2.Name, address2.Phone) //李四 155555555
}
