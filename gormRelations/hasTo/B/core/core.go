package core

import (
	"fmt"
	"golangproject01/gormRelations/hasTo/B/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	//这里需要先创建users表，再创建credit_cards表
	DB.AutoMigrate(&models.User{}, &models.CreditCard{})
}
