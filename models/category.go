package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"primaryKey;unique"`
	Items  []Item
	UserID uint
}
