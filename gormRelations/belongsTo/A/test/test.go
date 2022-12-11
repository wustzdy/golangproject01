package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/belongsTo/A/core"
	"golangproject01/gormRelations/belongsTo/A/models"
)

func main() {
	Init()
	createTest()
	selectTest()

	var map1 = map[string]string{"workspaceID": "111"}
	json, _ := json.Marshal(map1)
	fmt.Println("json:", string(json))
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
