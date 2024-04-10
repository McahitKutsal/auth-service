package services

import (
	"github.com/McahitKutsal/auth-service/base"
	"github.com/McahitKutsal/auth-service/models"
	"github.com/McahitKutsal/auth-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
	userValidator  base.UserService
}

func NewUserService(repo repositories.UserRepository, validator base.UserService) base.UserService {
	return &userServiceImpl{
		userRepository: repo,
		userValidator:  validator,
	}
}

func (s *userServiceImpl) Register(email, password string) error {

	// Validate input
	err := s.userValidator.Register(email, password)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	//role := models.Role{Name: "ADMIN", AccessRights: []models.AccessRight{models.AccessRight{Name: "read"}}}

	user := &models.User{Email: email, Password: string(hash)}
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Authenticate(email, password string) (*models.User, error) {

	_, err := s.userValidator.Authenticate(email, password)
	if err != nil {
		return nil, err
	}

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
