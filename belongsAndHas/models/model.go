package models

import (
	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
}
