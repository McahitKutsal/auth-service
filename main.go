package main

import (
	"os"

	"github.com/McahitKutsal/auth-service/controllers"
	"github.com/McahitKutsal/auth-service/initializers"
	"github.com/McahitKutsal/auth-service/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/"+os.Getenv("VERSION")+"/signup", controllers.SignUp)
	r.POST("/"+os.Getenv("VERSION")+"/login", controllers.Login)
	r.GET("/"+os.Getenv("VERSION")+"/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run()
}
