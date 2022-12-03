package main

import (
	"golangproject01/belongsTo/core"
	"golangproject01/belongsTo/models"
	"gorm.io/gorm"
)

func main() {
	initCore()
	oneToOne()
}

func initCore() {
	core.Init()
}

func oneToOne() {

	g := models.GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "美女",
	}

	d := models.Dog{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    "张三",
		GirlGod: g,
	}
	core.DB.Create(&d)
}
