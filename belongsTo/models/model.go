package models

import (
	"gorm.io/gorm"
)

// belongs to
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
