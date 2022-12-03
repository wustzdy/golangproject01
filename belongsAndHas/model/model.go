package model

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	name      string
	GirlGodID uint
	GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
}

//备注
