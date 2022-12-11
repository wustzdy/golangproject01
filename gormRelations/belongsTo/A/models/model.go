package models

import "gorm.io/gorm"

// `User` 属于 `Company`，`CompanyID` 是外键
// user 是属于 company 的，就是每个 user 有且只能对应分配给一个 company。
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}
