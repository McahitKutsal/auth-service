// middlewares/authMiddleware.go
package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/McahitKutsal/auth-service/repositories"
	"github.com/McahitKutsal/auth-service/utils"
	"github.com/gin-gonic/gin"
)

var userRepository = repositories.NewUserRepository()

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	requiredAccessRight := c.GetHeader("Required-Access-Right")
	fmt.Println(requiredAccessRight)
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := utils.ParseToken(tokenString)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	claims, ok := utils.ExtractClaims(token)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	if float64(utils.GetCurrentTimeUnix()) > claims["exp"].(float64) {
		utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	// Fetch user details from database based on userID
	// user, err := userRepository.FindUserByID(int(claims["id"].(float64)))
	// Handle the user retrieval and response here...
	userID := int(claims["id"].(float64))
	user, err := userRepository.FindUserByID(userID)
	if err != nil || user.ID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}
	c.Set("user", user)
	c.Next()
}
