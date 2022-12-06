package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         *string `gorm:"default:北京"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
}
