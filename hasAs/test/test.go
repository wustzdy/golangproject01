package main

import (
	"fmt"
	"golangproject01/hasAs/core"
	"golangproject01/hasAs/models"
)

func main() {
	initCore()
	//oneToOne()
	selectOne()
}

func initCore() {
	core.Init()
}

func hasAs() {

	d := models.Dog{
		Name: "张三",
	}
	g := models.GirlGod{
		Name: "美女",
		Dog:  d,
	}
	core.DB.Create(&g)
}

func selectOne() {
	var girl models.GirlGod
	core.DB.First(&girl, 1)
	fmt.Println("girl:", girl)

	//这个时候 只会查询girl相关的，但是dog查不出来， 需要改造
	var girl1 models.GirlGod
	core.DB.Preload("Dog").First(&girl1, 1)
	fmt.Println("girl:", girl1)

}
