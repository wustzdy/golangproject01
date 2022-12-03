package model

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	name string
}

type GirlGod struct {
	gorm.Model
	name string
}
