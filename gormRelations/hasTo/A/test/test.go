package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/hasTo/A/core"
	"golangproject01/gormRelations/hasTo/A/models"
)

func main() {
	Init()
	//createTest()
	selectTest()
	selectSql()
	selectSql1()
}

func Init() {
	core.Init()
}
func createTest() {

	u := models.User{
		CreditCard: models.CreditCard{
			Number: "dddd",
		},
	}
	core.DB.Create(&u)

}
func selectTest() {
	var user []models.User
	core.DB.Model(models.User{}).Preload("CreditCard").Find(&user)
	json, _ := json.Marshal(user)
	fmt.Println("user:", string(json))

}
func selectSql() (list []models.User, err error) {
	var sql = "select u.* from users u "
	err = core.DB.Raw(sql).Scan(&list).Error
	result, _ := json.Marshal(list)
	fmt.Println("result:", string(result))
	return

}

func selectSql1() (list []models.Result, err error) {
	var sql = "select u.id,u.created_at,u.updated_at,c.number,c.user_id from users u,credit_cards c where u.id=c.user_id"
	err = core.DB.Raw(sql).Scan(&list).Error
	result, _ := json.Marshal(list)
	fmt.Println("result:", string(result))
	return

}
