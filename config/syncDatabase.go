package config

import "github.com/McahitKutsal/auth-service/models"

func SyncDatabase() {
	// Auto Migrate
	// Auto Migrate
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.AccessRight{})
	DB.AutoMigrate(&models.UserRole{})
	DB.AutoMigrate(&models.RoleAccessRight{})
}
