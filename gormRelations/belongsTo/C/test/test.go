package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/belongsTo/C/core"
	"golangproject01/gormRelations/belongsTo/C/models"
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
	user := models.User{
		Name: "张三",
		Company: models.Company{
			Name: "上海科技",
		},
	}
	core.DB.Create(&user)
}

func selectTest() {
	var user models.User
	core.DB.Find(&user)
	b, _ := json.Marshal(user)
	fmt.Println("user:", string(b))
}
