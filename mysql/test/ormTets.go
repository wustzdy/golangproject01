package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID      uint
	Name    string
	Gender  string
	Hobby   string
	Address string
}

// 定义模型
type Student struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:age_of_the_beast_111"`
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`
	TypeScene    string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println(err)
	}

	createUserInfo(db)
	createStudent(db)
}

func createUserInfo(db *gorm.DB) {
	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{5, "七米", "男", "篮球", "北京"}
	u2 := UserInfo{6, "沙河娜扎", "女", "足球", "上海"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	/*// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)*/
	isExistField := db.Migrator().HasColumn(&UserInfo{}, "Address")
	if !isExistField {
		err := db.Migrator().AddColumn(&UserInfo{}, "Address")
		if err != nil {
			fmt.Printf("添加字段错误,err:%s\n", err)
			//return
		}
	}
	fmt.Printf("11111")
}

func createStudent(db *gorm.DB) {
	//db.AutoMigrate(&Student{})

	isExistField := db.Migrator().HasColumn(&Student{}, "TypeScene")
	if !isExistField {
		err := db.Migrator().AddColumn(&Student{}, "TypeScene")
		if err != nil {
			fmt.Printf("添加字段错误,err:%s\n", err)
			//return
		}
	}
	fmt.Printf("11111")
}
