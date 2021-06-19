package models

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	gorm.Model
	Content    string `gorm:"size:65535" json:"content"`
	CreateTime int64  `gorm:"column:creattime" json:"build"`
	Role       string `json:"role"`
}
