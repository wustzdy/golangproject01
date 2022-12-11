package models

import "gorm.io/gorm"

// 例如我们想自定义外键，就需要用标签foreignKe来指定外键：
type User struct {
	gorm.Model
	Name      string
	CompanyID string
	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
}

type Company struct {
	ID   int
	Code string
	Name string
}
