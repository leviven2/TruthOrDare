package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string
	Items  []Item
	UserID uint
}
