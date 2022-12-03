package models

import (
	"gorm.io/gorm"
)

// has to
type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
}

type GirlGod struct {
	gorm.Model
	Name string
	Dog  Dog
}
