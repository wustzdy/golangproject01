package main

import "fmt"

func TestUpdate() {
	//update 只更新你选择的字段
	//updates 更新所有字段，有两种形式，一种是map,一种是结构体 结构体零值不参与更新
	//save 无论如何都更新所有内容
	//把查出来的name模糊匹配到到的武汉，name更新为朱大洋
	DB.Model(&TestUser{}).Where("name like ?", "%武汉%").Update("name", "朱大洋")
}
func TestUpdateSave() {
	var users []TestUser
	dbRes := DB.Model(&TestUser{}).Where("name like ?", "%朱%").Find(&users)
	for k := range users {
		users[k].Age = 18
	}
	dbRes.Save(&users)

}
func TestUpdates() {
	var user TestUser
	//查询第一条
	DB.First(&user).Updates(&TestUser{
		Name: "上海交通大学",
		Age:  0,
	})
	fmt.Println("userFirst:", user) //可以看到零值和空字符串不参与更新，这时候可以采用map

	var user1 TestUser
	DB.First(&user1).Updates(map[string]interface{}{
		"name": " ",
		"age":  0,
	})

}
