package models

type User struct {
	Id       int
	UserName string
	Age      int
	Email    string
}

func (User) TableName() string {
	return "users"
}
