// utils/responseUtils.go
package utils

import (
	"net/http"

	"github.com/McahitKutsal/auth-service/models"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	response := models.Response{
		Status: http.StatusOK,
		Data:   data,
		Errors: []string{},
	}
	c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, status int, err error) {
	response := models.Response{
		Status: status,
		Data:   nil,
		Errors: []string{err.Error()},
	}
	c.JSON(status, response)
}

func CustomResponse(c *gin.Context, status int, data interface{}, errors []string) {
	response := models.Response{
		Status: status,
		Data:   data,
		Errors: errors,
	}
	c.JSON(status, response)
}
