package models

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Quote struct {
	gorm.Model
	Client        string `gorm:"column:client"`
	Telephone     string `gorm:"column:telephone"`
	DateQuoteInit string `gorm:"column:dateQuoteInit"`
	DateQuoteEnd  string `gorm:"column:dateQuoteEnd"`
	UserID        uint
	IsDel         soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
