package main

import (
	"golangproject01/hasAs/core"
	"golangproject01/hasAs/models"
)

func main() {
	initCore()
	oneToOne()
}

func initCore() {
	core.Init()
}

func oneToOne() {

	d := models.Dog{
		Name: "张三",
	}
	g := models.GirlGod{
		Name: "美女",
		Dog:  d,
	}
	core.DB.Create(&g)
}
