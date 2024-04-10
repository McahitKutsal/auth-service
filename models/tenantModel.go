package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Users []User `gorm:"foreignKey:TenantID"`
	//title will be added
}
