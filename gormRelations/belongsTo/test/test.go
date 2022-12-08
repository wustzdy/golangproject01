package main

import (
	"golangproject01/gormRelations/belongsTo/core"
	"golangproject01/gormRelations/belongsTo/models"
)

func main() {
	Init()
	createTest()
}

func Init() {
	core.Init()
}

func createTest() {
	user := models.User{
		Name: "张三",
		Company: models.Company{
			Name: "上海科技",
		},
	}
	core.DB.Create(&user)
}
