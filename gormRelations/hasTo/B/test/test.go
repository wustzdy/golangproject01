package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/hasTo/B/core"
	"golangproject01/gormRelations/hasTo/B/models"
)

func main() {
	Init()
	createTest()
	selectTest()
}

func Init() {
	core.Init()
}
func createTest() {

	u := models.User{
		CreditCard: models.CreditCard{
			Number:   "dddd",
			UserName: "张三",
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
