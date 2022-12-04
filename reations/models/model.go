package models

import (
	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	Info      Info
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog
}

type Info struct {
	gorm.Model
	Money int
	DogID uint
}
