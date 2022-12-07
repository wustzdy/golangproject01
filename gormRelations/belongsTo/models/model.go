package models

import "gorm.io/gorm"

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
	//// 使用 CompanyRefer 作为外键
}

type Company struct {
	ID   int
	Name string
}
