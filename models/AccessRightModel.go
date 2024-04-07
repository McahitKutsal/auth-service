package models

import "gorm.io/gorm"

type AccessRight struct {
	gorm.Model
	Name string `gorm:"unique"`
}
