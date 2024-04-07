package routes

import (
	"os"

	"github.com/McahitKutsal/auth-service/controllers"
	_ "github.com/McahitKutsal/auth-service/docs"
	"github.com/McahitKutsal/auth-service/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	v := os.Getenv("VERSION")

	// Swagger documentation route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	authGroup := r.Group("/" + v + "/api")
	{
		authGroup.POST("/signup", controllers.SignUp)
		authGroup.POST("/login", controllers.Login)
		authGroup.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	}
}
