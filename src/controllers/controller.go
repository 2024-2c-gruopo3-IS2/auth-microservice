// controllers/auth_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"auth-microservice/services"
)

func SignupHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		IsAdmin  bool   `json:"is_admin"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.RegisterUser(req.Email, req.Password, req.IsAdmin)
	if err != nil {
		if err.Error() == "User already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := services.LoginUser(req.Email, req.Password, req.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func SigninHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		IsAdmin  bool   `json:"is_admin"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.LoginUser(req.Email, req.Password, req.IsAdmin)
	if err != nil {
		if err.Error() == "user is blocked" {
			c.JSON(http.StatusForbidden, gin.H{"error": "user is blocked"})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func GetEmailFromTokenHandler(c *gin.Context) {
    var req struct {
        Token string `json:"token" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    email, err := services.GetEmailFromToken(req.Token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"email": email})
}

func BlockUserHandler(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.BlockUser(req.Email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully blocked"})
}