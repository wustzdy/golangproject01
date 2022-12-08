package main

import (
	"fmt"
	"golangproject01/gormRelations/hasTo/core"
	"golangproject01/gormRelations/hasTo/models"
)

func main() {
	Init()
	//createTest()
	selectTest()
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
	var user models.User
	core.DB.Model(models.User{}).Preload("CreditCard").Find(&user)
	fmt.Println("user:", user)

}
