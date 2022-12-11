package models

import (
	"gorm.io/gorm"
)

// CreditCard表里的 UserName 关联到User表里的 name字段上
type User struct {
	gorm.Model
	MemberNumber string     `gorm:"index"`
	CreditCard   CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}
