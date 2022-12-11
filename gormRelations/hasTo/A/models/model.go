package models

import (
	"gorm.io/gorm"
)

// User 有一张 CreditCard，UserID 是外键
// 1，user 是拥有 creditcard 的，creditcard 有且只能被一个 user 拥有。
// 2，Creditcard是属于user的，或者说user关联Creditcard
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
