package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/hasTo/C/core"
	"golangproject01/gormRelations/hasTo/C/models"
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
			Number: "123",
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
