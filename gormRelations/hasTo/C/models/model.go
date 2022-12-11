package models

import (
	"gorm.io/gorm"
)

// CreditCard表里的 UserName 关联到User表里的 name字段上
type User struct {
	gorm.Model
	Name       string     `sql:"index"`
	CreditCard CreditCard `gorm:"foreignkey:UserName;references:Name"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}
