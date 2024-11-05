// main.go
package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"auth-microservice/config"
	"auth-microservice/controllers"
)

func main() {
	config.InitDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.SignupHandler)
		authRoutes.POST("/signin", controllers.SigninHandler)
		authRoutes.GET("/get-email-from-token", controllers.GetEmailFromTokenHandler)
		authRoutes.POST("/block-user", controllers.BlockUserHandler)
		authRoutes.POST("/unblock-user", controllers.UnblockUserHandler)
		authRoutes.GET("/get-users-status", controllers.GetUsersStatusHandler)
		authRoutes.POST("/request-password-reset", controllers.RequestPasswordResetHandler)
		authRoutes.POST("/password-reset", controllers.ResetPasswordHandler)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if err := router.Run(host + ":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
