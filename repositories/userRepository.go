// userRepository.go
package repositories

import "github.com/McahitKutsal/auth-service/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id int) (*models.User, error)
}
