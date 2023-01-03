package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"type:varchar(32) NOT NULL DEFAULT ''"`
	Description string `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
}

type Meta struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(32) NOT NULL DEFAULT ''"`
	Email     string `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	ProjectId uint   `gorm:"type:bigint unsigned NOT NULL DEFAULT 0"`
}

type ProjectMetas struct {
	ProjectId   uint   `json:"projectId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UserName    string `json:"userName,omitempty"`
	Email       string `json:"email,omitempty"`
}
