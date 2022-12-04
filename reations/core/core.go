package core

import (
	"fmt"
	"golangproject01/reations/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	DB.AutoMigrate(&models.Dog{})
	DB.AutoMigrate(&models.GirlGod{})
	DB.AutoMigrate(&models.Info{})
}
