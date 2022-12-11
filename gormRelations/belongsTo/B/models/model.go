package models

import "gorm.io/gorm"

// 例如我们想自定义外键，就需要用标签foreignKe来指定外键：
type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
	// 使用 CompanyRefer 作为外键
}

type Company struct {
	ID   int
	Name string
}
