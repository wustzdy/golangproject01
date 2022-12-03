package main

import (
	"fmt"
	"golangproject01/belongsAndHas/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	DB.AutoMigrate(&model.Dog{}) //会自动创建关联表
}
