package main

import (
	"golangproject01/gormRelations/hasTo/core"
	"golangproject01/gormRelations/hasTo/models"
)

func main() {
	Init()
	createTest()
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
