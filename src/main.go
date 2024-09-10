// main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"auth-microservice/config"
	"auth-microservice/controllers"
)

func main() {
	config.InitDB()

	router := gin.Default()

	authRoutes := router.Group("/auth")

	{
		authRoutes.POST("/signup", controllers.SignupHandler)
		authRoutes.POST("/signin", controllers.SigninHandler)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if err := router.Run(host + ":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}