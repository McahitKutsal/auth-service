// validators/userValidator.go
package validators

import (
	"github.com/McahitKutsal/auth-service/base"
	"github.com/McahitKutsal/auth-service/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type userValidatorImpl struct{}

func NewUserValidator() base.UserService {
	return &userValidatorImpl{}
}

func (s *userValidatorImpl) Register(email, password string) error {

	err := validate.Var(email, "required,email")
	if err != nil {
		return err
	}

	err = validate.Var(password, "required,min=8")
	if err != nil {
		return err
	}

	return nil
}

func (s *userValidatorImpl) Authenticate(email, password string) (*models.User, error) {
	err := validate.Var(email, "required,email")
	if err != nil {
		return nil, err
	}

	err = validate.Var(password, "required,min=8")
	if err != nil {
		return nil, err
	}

	return nil, nil
}
