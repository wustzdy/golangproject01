package main

import (
	"database/sql"
	"fmt"
	"golangproject01/gormRelations/core"
	"golangproject01/gormRelations/models"
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
	core.DB.Find(&user)
	fmt.Printf("name:%#v,email:%#v,age:%#v,memberNumber:%#v:",
		*user.Name, *user.Email, int(user.Age), user.MemberNumber.String)

	fmt.Printf("user111:", user)

}
