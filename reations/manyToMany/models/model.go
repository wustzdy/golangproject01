package models

import (
	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name     string
	Info     Info
	GirlGods []GirlGod `gorm:"many2many:dog_girl_god"` //中间表
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god"`
}

type Info struct {
	gorm.Model
	Money int
	DogID uint
}
