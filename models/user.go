package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"primaryKey;unique"`
	Password   string
	Categories []Category
}
