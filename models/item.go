package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Description string
	CategoryID  uint
}
