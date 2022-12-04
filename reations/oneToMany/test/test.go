package main

import (
	"fmt"
	"golangproject01/reations/oneToMany/core"
	"golangproject01/reations/oneToMany/models"
	"gorm.io/gorm"
)

func main() {
	initCore()
	//createTest()
	selectHasMany()
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

	//在Dog里面加入info
	var girl3 models.GirlGod
	core.DB.Preload("Dogs.Info").Preload("Dogs").First(&girl3)
	fmt.Println("girl3:", girl3)

	//在Dog里面加入info
	var girl4 models.GirlGod
	core.DB.Preload("Dogs.Info").First(&girl4)
	fmt.Println("girl4:", girl4)

	//如果是过滤的话
	//在Dog里面加入info
	var girl5 models.GirlGod
	core.DB.Preload("Dogs.Info", "money>100").Preload("Dogs", "name=?", "张三").First(&girl5)
	fmt.Println("girl4:", girl5)

	//查询大于200money的那个人
	var girl6 models.GirlGod
	core.DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money>100")
	}).First(&girl6)
	fmt.Println("girl6:", girl6)
}
