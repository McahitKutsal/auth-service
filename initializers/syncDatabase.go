package initializers

import "github.com/McahitKutsal/auth-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
