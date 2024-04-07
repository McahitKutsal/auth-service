// services/userService.go
package services

import (
	"github.com/McahitKutsal/auth-service/models"
)

type UserService interface {
	Register(email, password string) error
	Authenticate(email, password string) (*models.User, error)
}
