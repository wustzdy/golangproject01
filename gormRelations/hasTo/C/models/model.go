package models

import (
	"gorm.io/gorm"
)

// 如果你想要使用另一个字段来保存该关系，你同样可以使用标签 foreignKey 来更改它，例如：
type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"foreignkey:user_name;association_foreignkey:name"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}
