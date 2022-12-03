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

	var user1 TestUser
	DB.Where("name=?", "张三").Find(&user1)
	fmt.Println("selectUser1:", user1)

	var user2 TestUser
	DB.Where(&TestUser{
		Name: "张三",
	}).Find(&user2)
	fmt.Println("selectUser2:", user2)

	var user3 TestUser
	DB.Where(map[string]interface{}{
		"name": "张三",
	}).First(&user3)
	fmt.Println("selectUser3:", user3)

	//查询主键为6的记录
	var user4 TestUser
	DB.First(&user4, "6")
	fmt.Println("selectUser4:", user4)

	//内联查询-string形式
	var user5 TestUser
	DB.First(&user5, "name=?", "张三")
	fmt.Println("selectUser5:", user5)

	//内联查询-map形式
	var user6 TestUser
	DB.First(&user6, map[string]interface{}{
		"name": "张三",
	})
	fmt.Println("selectUser6:", user6)

	//内联查询-对象形式
	var user7 TestUser
	DB.First(&user7, TestUser{
		Name: "武汉科技大学",
	})
	fmt.Println("selectUser7:", user7)

	//查询多条记录
	var user8 []TestUser
	DB.Find(&user8)
	fmt.Println("selectUser8:", user8)

	//查询多条记录-筛选
	var user9 []TestUser
	DB.Where("name <> ?", "张三").Find(&user9)
	fmt.Println("selectUser9:", user9)

	//查询多条记录-模糊匹配
	var user10 []TestUser
	DB.Where("name like ?", "%武汉%").Find(&user10)
	fmt.Println("selectUser10:", user10)

	//查询多条记录-模糊匹配-只需要name返回字段
	var user11 []TestUser
	DB.Select("name").Where("name like ?", "%武汉%").Find(&user11)
	fmt.Println("selectUser12:", user11)
}
