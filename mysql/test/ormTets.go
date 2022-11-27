package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	db.AutoMigrate(&UserInfo{})

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "七米", "男", "篮球"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
