package models

type SceneSource int

const (
	TaskAutoInner  SceneSource = 1
	TaskAutoDefine SceneSource = 2
)

type User struct {
	Id          int
	UserName    string
	Age         int
	Email       string
	name        string
	SceneSource SceneSource `gorm:"type:tinyint(1) NOT NULL DEFAULT 1;index"`
}

func (User) TableName() string {
	return "user"
}
