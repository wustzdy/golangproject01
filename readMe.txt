https://blog.csdn.net/weixin_46618592/article/details/127173491?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EYuanLiJiHua%7EPosition-2-127173491-blog-124781950.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EYuanLiJiHua%7EPosition-2-127173491-blog-124781950.pc_relevant_default&utm_relevant_index=3

3、数据表字段接口

type User struct {
	Sex bool
}
err := db.Migrator().AddColumn(&User{}, "Sex")
if err != nil {
	fmt.Printf("添加字段错误,err:%s\n", err)
	return
}

2）DropColumn 删除字段
err := db.Migrator().DropColumn(&User{}, "Age")
if err != nil {
	fmt.Printf("删除字段错误,err:%s\n", err)
	return
}
3）RenameColumn 修改字段名
注意：

    必须先声明模型。
    修改的字段名在对应的数据表必须存在，修改的字段名和修改后的字段名必须定义在结构体内。

type User struct {
	Name     string
	UserName string
}
err := db.Migrator().RenameColumn(&User{}, "name", "user_name")
if err != nil {
	fmt.Printf("修改字段名错误,err:%s\n", err)
	return
}
4）HasColumn 查询字段是否存在
isExistField := db.Migrator().HasColumn(&User{}, "name")
fmt.Printf("name字段是否存在:%t\n", isExistField)
isExistField = db.Migrator().HasColumn(&User{}, "user_name")
fmt.Printf("user_name:%t\n", isExistField)
4、 数据库表的索引接口
1）CreateIndex 为字段创建索引
注意：

    必须先声明模型。
    必须先在声明模型中使用标签gorm:index定义索引。

type User struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}
// 为 Name 字段创建索引,两种方法都可以
db.Migrator().CreateIndex(&User{}, "Name")
db.Migrator().CreateIndex(&User{}, "idx_name")

2）DropIndex 为字段删除索引
db.Migrator().DropIndex(&User{}, "Name")
db.Migrator().DropIndex(&User{}, "idx_name")

）HasIndex 检查索引是否存在

isExists := db.Migrator().HasIndex(&User{}, "idx_name")
fmt.Printf("idex_name是否存在:%t\n", isExists)

db.Migrator().CreateIndex(&User{}, "idx_name")
isExists = db.Migrator().HasIndex(&User{}, "idx_name")
fmt.Printf("idex_name是否存在:%t\n", isExists)
————————————————
版权声明：本文为CSDN博主「lin钟一」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_46618592/article/details/127173491

————————————————
版权声明：本文为CSDN博主「lin钟一」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_46618592/article/details/127173491
————————————————
版权声明：本文为CSDN博主「lin钟一」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_46618592/article/details/127173491
