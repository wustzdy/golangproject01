package main

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:my_time"`
}
type TestUser struct {
	gorm.Model
	Name string `gorm:"default:qm"`
	Age  uint8  `gorm:"comment:年纪"`
}

type UserInfo struct {
	Name string
	Age  uint8
}

func TestUserCreate() {
	DB.AutoMigrate(&TestUser{})
}
