package main

import (
	"database/sql"
	"fmt"
	"golangproject01/gormRelations/base/core"
	"golangproject01/gormRelations/base/models"
)

func main() {
	Init()
	//CreateTest()
	selectTest()
}

func Init() {
	core.Init()
}
func CreateTest() {
	var user1 = models.User{
		Email:        new(string),
		Age:          0,
		MemberNumber: sql.NullString{String: "", Valid: true},
	}
	core.DB.Create(&user1) /// 此时数据库中该条记录name字段的值就是 默认值 北京

	var user2 = models.User{
		Name:         new(string),
		Email:        new(string),
		Age:          0,
		MemberNumber: sql.NullString{String: "", Valid: true},
	}
	core.DB.Create(&user2) /// 此时数据库中该条记录name字段的值就是''，传空

	var name = "上海"
	var user3 = models.User{
		Name:         &name,
		Email:        new(string),
		Age:          0,
		MemberNumber: sql.NullString{String: "wustzdy", Valid: true},
	}
	core.DB.Create(&user3) //上海
}

func selectTest() {
	var user models.User
	core.DB.Last(&user)

	fmt.Printf("name:%#v,email:%#v,age:%#v,memberNumber:%#v:",
		*user.Name, *user.Email, int(user.Age), user.MemberNumber.String)

	core.DB.Where("name=?", "上海").Find(&user)
	fmt.Println("user1:", user)

	var users []models.User
	core.DB.Where("name <> ?", "上海").Find(&users)
	fmt.Println("users:", users)

	var users1 []models.User
	core.DB.Where("name in (?)", []string{"上海", "北京"}).Find(&users1)
	fmt.Println("users1:", users1)

	//struct查询
	var user1 models.User
	var name = "北京"
	core.DB.Where(&models.User{
		Name: &name,
		Age:  20,
	}).First(&user1)
	fmt.Println("userStruct:", user1)

	//map查询
	var user2 []models.User
	core.DB.Where(map[string]interface{}{
		"Name": "北京",
		"Age":  0,
	}).Find(&user2)
	fmt.Println("mapStruct:", user2)

	//or
	var user3 []models.User
	core.DB.Where("name=?", "上海").Or("name=?", "北京").Find(&user3)
	fmt.Println("user3:", user3)

	//struct or
	var user4 []models.User
	var name1 = "北京"
	core.DB.Where("name=?", "上海").Or(&models.User{
		Name: &name1,
	}).Find(&user4)
	fmt.Println("user4:", user4)

	//map or
	var user5 []models.User
	core.DB.Where("name=?", "上海").Or(map[string]interface{}{
		"name": "上海",
		"age":  0,
	}).Find(&user5)
	fmt.Println("user5:", user5)

}
