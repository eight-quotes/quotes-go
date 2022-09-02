package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"column:name"`
	UserName string  `gorm:"not null;unique_index"`
	Password string  `gorm:"column:password" json:"-"`
	Quotes   []Quote `json:"quotes"`
}
