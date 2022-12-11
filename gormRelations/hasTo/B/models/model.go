package models

import (
	"gorm.io/gorm"
)

// 如果你想要使用另一个字段来保存该关系，你同样可以使用标签 foreignKey 来更改它，例如：
type User struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignKey:UserName"`
	// use UserName as foreign key
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

/*可以看到，CreditCard表最后的字段确实是变成了user_name，但是user_name还是1 ，
也就是说，还是把User的ID字段直接拿过来放在了里面，我们可以再确认一下，结构体定义不变，看下方代码*/
