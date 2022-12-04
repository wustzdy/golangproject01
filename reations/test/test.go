package main

import (
	"fmt"
	"golangproject01/reations/core"
	"golangproject01/reations/models"
	"gorm.io/gorm"
)

func main() {
	initCore()
	//createTest()
	//selectHasMany()
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

func selectHasMany() {
	var girl models.GirlGod
	core.DB.Preload("Dogs").First(&girl) //得到两条
	fmt.Println("girl:", girl)

	//得到一条
	var girl1 models.GirlGod
	core.DB.Preload("Dogs", "name=?", "张三").First(&girl1) //得到两条
	fmt.Println("girl1:", girl1)

	//第二种写法
	var girl2 models.GirlGod
	core.DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("name", "张三")
	}).First(&girl2)
	fmt.Println("girl2:", girl2)
}
