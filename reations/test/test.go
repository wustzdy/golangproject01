package main

import (
	"golangproject01/reations/core"
	"golangproject01/reations/models"
	"gorm.io/gorm"
)

func main() {
	initCore()
	createTest()
}

func initCore() {
	core.Init()
}

func createTest() {
	d1 := models.Dog{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "张三",
	}
	d2 := models.Dog{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "李四",
	}
	g := models.GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "美女",
		Dogs: []models.Dog{d1, d2},
	}
	core.DB.Create(&g)

}
