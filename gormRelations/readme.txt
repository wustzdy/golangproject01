7.所有字段的零值都不会被保存到数据库，但会自动填充预设的默认值。===>

（1）如果想更新零值，在定义结构体字段时，给指针类型，如

（2）或者给 sql.NullString 类型


1,
type UserInfo struct {
    Name  *string  `json:"name"`
}
DB.Create(&UserInfo{ Name: new(string) }) // 这里需要构建一个 new(string) 字符串指针，不然传入类型不匹配。
2,
type UserInfo struct {
   Name  sql.NullString `json:"name"`
}
DB.Create(&UserInfo{ Name: sql.NullString{String:"", Vaild：true} }) // 这里需要构建一个 new(string) 字符串指针，不然传入类型不匹


3，
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}Copy to clipboardErrorCopied


4，
type User struct {} // 默认表名是 `users`

// 将 User 的表名设置为 `profiles`
5，
func (User) TableName() string {
  return "profiles"
}

func (u User) TableName() string {
  if u.Role == "admin" {
    return "admin_users"
  } else {
    return "users"
  }
}

6，
// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
db.SingularTable(true)Copy to clipboardErrorCopied


7，
也可以通过Table()指定表名：
// 使用User结构体创建名为`deleted_users`的表
db.Table("deleted_users").CreateTable(&User{})

var deleted_users []User
db.Table("deleted_users").Find(&deleted_users)
 SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
 DELETE FROM deleted_users WHERE name = 'jinzhu';Copy to clipboardErrorCopied
8，
可以使用结构体tag指定列名：
type Animal struct {
  AnimalId    int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
  Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
  Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}Copy to clipboardErrorCopied

9，时间戳跟踪
CreatedAt
如果模型有 CreatedAt字段，该字段的值将会是初次创建记录的时间。

db.Create(&user) // `CreatedAt`将会是当前时间

// 可以使用`Update`方法来改变`CreateAt`的值
db.Model(&user).Update("CreatedAt", time.Now())

10，
UpdatedAt
如果模型有UpdatedAt字段，该字段的值将会是每次更新记录的时间。
db.Save(&user) // `UpdatedAt`将会是当前时间
db.Model(&user).Update("name", "jinzhu") // `UpdatedAt`将会是当前时间

11，
如果模型有DeletedAt字段，调用Delete删除该记录时，将会设置DeletedAt字段为当前时间，而不是直接将记录从数据库中删除。

12，
创建-创建记录
首先定义模型：
type User struct {
    ID           int64
    Name         string
    Age          int64
}
使用使用NewRecord()查询主键是否存在，主键为空使用Create()创建记录：
user := User{Name: "q1mi", Age: 18}

db.NewRecord(user) // 主键为空返回`true`
db.Create(&user)   // 创建user
db.NewRecord(user) // 创建`user`后返回`false`

13，默认值
可以通过 tag 定义字段的默认值，比如：

type User struct {
  ID   int64
  Name string `gorm:"default:'小王子'"`
  Age  int64
}
**注意：**通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。

举个例子：

var user = User{Name: "", Age: 99}
db.Create(&user)
上面代码实际执行的SQL语句是INSERT INTO users("age") values('99');，排除了零值字段Name，
而在数据库中这一条数据会使用设置的默认值小王子作为Name字段的值。

**注意：**所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
 如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口，比如：

使用指针方式实现零值存入数据库
// 使用指针
type User struct {
  ID   int64
  Name *string `gorm:"default:'小王子'"`
  Age  int64
}
user := User{Name: new(string), Age: 18))}
db.Create(&user)  // 此时数据库中该条记录name字段的值就是''

使用Scanner/Valuer接口方式实现零值存入数据库
// 使用 Scanner/Valuer
type User struct {
    ID int64
    Name sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
    Age  int64
}
user := User{Name: sql.NullString{"", true}, Age:18}
db.Create(&user)  // 此时数据库中该条记录name字段的值就是''

type User struct {
	gorm.Model
	Name         *string `gorm:"default:北京"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
}

func selectTest() {
	var user models.User
	core.DB.Last(&user)
	fmt.Printf("name:%#v,email:%#v,age:%#v,memberNumber:%#v:",
		*user.Name, *user.Email, int(user.Age), user.MemberNumber.String)
}
