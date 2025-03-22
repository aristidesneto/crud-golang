package main

import (
	"log"

	"github.com/aristidesneto/crud-golang/src/config/logger"
	"github.com/aristidesneto/crud-golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup)

	// log.Println("Aplcation is running...")
	logger.Info("Application is running...")
	logger.Debug("Iam a debug log")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
