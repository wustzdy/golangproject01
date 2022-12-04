package main

import (
	"fmt"
	"golangproject01/reations/defineData/core"
	"golangproject01/reations/defineData/models"
)

func main() {
	initCore()
	//createTest()
	selectTest()
}

func initCore() {
	core.Init()
}

func createTest() {
	cuser := models.CUser{
		CInfo: models.CInfo{
			Name: "张三",
			Age:  13,
		},
		ID: 1,
	}
	core.DB.Create(&cuser)
}

func selectTest() {

	var cuser models.CUser
	core.DB.Find(&cuser)
	fmt.Println("cuser:", cuser)
}
