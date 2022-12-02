package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表'").AutoMigrate(&Posts{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户'").AutoMigrate(&Users{})
}

type Users struct {
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"unique_index"`
	// 一对多关系映射，一个用户有多篇文章，用户表的id作为关联键
	PostsArticle []Posts `gorm:"FOREIGNKEY:Userid;ASSOCIATION_FOREIGNKEY:ID"`
}
type Posts struct {
	ID uint `gorm:"primarykey"`
	//Key string `gorm:"unique:not null"`
	Userid     int
	Title      string `gorm:"type:varchar(200)"`
	Summart    string `gorm:"varchar(800)"`
	Content    string `gorm:"type:text"`
	VisitCount int    `gorm:"default:0"`
	Like       int    `gorm:"default:0"`
}

func main() {
	/*var user = new(Users)
	DB.Preload("PostsArticle").Where("id=1").Find(&user)*/
	/*user := &Users{
		ID:       1,
		Username: "h7",
		PostsArticle: []Posts{
			{
				ID:      1,
				Userid:  1,
				Title:   "这是一首简单的小情歌",
				Summart: "小情歌",
				Content: "001",
			},
			{
				ID:      2,
				Userid:  2,
				Title:   "东风破",
				Summart: "东风破001",
				Content: "002",
			},
		},
	}*/

	//DB.Preload("PostsArticle").Where("id=1").Find(&user)
	var user = new(Users)
	DB.Preload("PostsArticle").Where("id=1").Find(&user)
	fmt.Println(user, user.Username, user.PostsArticle)
	fmt.Println(user.PostsArticle[0].Title)
}
