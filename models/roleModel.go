package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name         string        `gorm:"unique"`
	AccessRights []AccessRight `gorm:"many2many:role_access_rights;"`
}
