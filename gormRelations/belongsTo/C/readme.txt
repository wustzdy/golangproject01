3）重写外键
要定义一个 belongs to 关系，数据库的表中必须存在外键。默认gorm使用（关联属性类型 + 主键）组成外键名，
如上面的例子User + ID 组成UserID，UserID就作为Profile的外键。

例如我们想自定义外键，就需要用标签foreignKe来指定外键：

type User struct {
  gorm.Model
  Name         string
  CompanyRefer int
  Company      Company `gorm:"foreignKey:CompanyRefer"`
  // 使用 CompanyRefer 作为外键
}

type Company struct {
  ID   int
  Name string
}