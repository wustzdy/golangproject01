package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 出差登记信息表
type Business struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	BusinessUsers []BusinessUser
}

// 出差人员表
type BusinessUser struct {
	gorm.Model
	UserID     int64    // 外键 (属于), tag `index`是为该列创建索引
	BusinessID uint     `json:"businessid"`        //这个对应business表中的ID
	NickNames  NickName `gorm:"foreignkey:UserID"` //加不加这个references:UserID没所谓，奇怪
}

// 按道理，上面应该是`gorm:"foreignkey:ID;references:UserID"`，即主表businessuser中的USERID=
// =从表中的ID啊

type NickName struct {
	gorm.Model
	NickName string
}

var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.AutoMigrate(&Business{}, &BusinessUser{}, &NickName{})
}
func main() {
	m := Business{}
	tx := DB.Find(&m)
	fmt.Println("query tx:", tx)

	/*	err = DB.Order("business.updated_at desc").
		Preload("BusinessUsers").Preload("BusinessUsers.NickNames").Where("business.project_id = ?", 1).
		Where("end_date > ?", time.Now()).
		Find(&Business{}).Error*/

}
