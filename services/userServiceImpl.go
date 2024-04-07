package services

import (
	"fmt"

	"github.com/McahitKutsal/auth-service/models"
	"github.com/McahitKutsal/auth-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: repo,
	}
}

func (s *userServiceImpl) Register(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	// rolesSlice := []*models.Role{
	// 	{Name: "Role1", AccessRights: []string{"read"}},
	// 	{Name: "Role2", AccessRights: []string{"write"}},
	// }
	// Create a new AccessRight
	// accessRight := models.AccessRight{Name: "read"}

	// // Create a new Role with the AccessRight
	// role := models.Role{Name: "ADMIN", AccessRights: []models.AccessRight{accessRight}}

	user := &models.User{Email: email, Password: string(hash)}

	fmt.Println("user", &user)
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil || user == nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
