package main

import "fmt"

func CreateTestUser() {
	tx := DB.Create(&TestUser{
		Name: "张三",
		Age:  18,
	})
	if tx != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}

// 只想创建name不想创建age
func CreateSingleTest() {
	DB.Select("Name").Create(&TestUser{
		Name: "李四",
		Age:  19,
	})
}

// 批量创建
func CreateBatch() {
	DB.Create(&[]TestUser{
		{Name: "武汉科技大学", Age: 128},
		{Name: "华中科技大学", Age: 128},
		{Name: "武汉大学", Age: 128},
		{Name: "武汉理工大学", Age: 128},
	})

}

func TestFind() {
	var result map[string]interface{}
	{
	}
	DB.Model(&TestUser{}).Find(&result)
	fmt.Println("result:", result)

	var result1 = make(map[string]interface{})
	DB.Model(&TestUser{}).Find(&result1)
	fmt.Println("result1:", result1)

	var user TestUser
	DB.Model(&TestUser{}).Find(&user)
	fmt.Println("user:", user)
}
