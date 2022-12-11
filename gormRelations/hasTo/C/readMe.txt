重写引用

默认情况下，拥有者实体会将 has one 对应模型的主键保存为外键，您也可以修改它，用另一个字段来保存，例如下个这个使用 Name 来保存的例子。

您可以使用标签 references 来更改它，例如：

CreditCard表里的 UserName 关联到User表里的 name字段上

CreditCard CreditCard gorm:"foreignkey:UserName;references:name"
type User struct {
  gorm.Model
  Name       string     `sql:"index"`
  CreditCard CreditCard `gorm:"foreignkey:UserName;references:name"`
}

type CreditCard struct {
  gorm.Model
  Number   string
  UserName string
}