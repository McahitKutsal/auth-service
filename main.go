package main

import (
	"github.com/McahitKutsal/auth-service/config"
	"github.com/McahitKutsal/auth-service/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDb()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}
