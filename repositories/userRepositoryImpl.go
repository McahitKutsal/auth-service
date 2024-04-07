// userRepositoryImpl.go
package repositories

import (
	"github.com/McahitKutsal/auth-service/config"
	"github.com/McahitKutsal/auth-service/models"
)

type userRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (*userRepositoryImpl) CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func (*userRepositoryImpl) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (*userRepositoryImpl) FindUserByID(id int) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
