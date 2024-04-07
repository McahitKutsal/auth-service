// controllers/userController.go
package controllers

import (
	"errors"
	"net/http"

	"github.com/McahitKutsal/auth-service/models"
	"github.com/McahitKutsal/auth-service/repositories"
	"github.com/McahitKutsal/auth-service/services"
	"github.com/McahitKutsal/auth-service/utils"
	"github.com/McahitKutsal/auth-service/validators"
	"github.com/gin-gonic/gin"
)

var userRepository = repositories.NewUserRepository()
var userValidator = validators.NewUserValidator()
var userService = services.NewUserService(userRepository, userValidator)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = userService.Register(body.Email, body.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "user have been created.")
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := userService.Authenticate(body.Email, body.Password)
	if err != nil || user == nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid username or password"))
		return
	}

	tokenString, err := utils.GenerateToken(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, gin.H{"token": tokenString})
}

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusBadRequest, nil)
		return
	}

	userToReturn := user.(*models.User)
	userToReturn.Password = ""
	utils.SuccessResponse(c, gin.H{"user": userToReturn})
}
