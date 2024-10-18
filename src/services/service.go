// services/auth_service.go
package services

import (
	"errors"
	"fmt"
	"auth-microservice/models"
	"auth-microservice/repositories"
	"auth-microservice/utils"
)

// RegisterUser handles registration logic for users and admins
func RegisterUser(email, password string, isAdmin bool) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = repositories.CreateUser(user, isAdmin)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser handles login logic for users and admins
func LoginUser(email, password string, isAdmin bool) (string, error) {
	if isAdmin {
		admin, err := repositories.GetAdminByEmail(email)
		
		if err != nil || !utils.CheckPasswordHash(password, admin.Password) {
			return "", errors.New("invalid credentials")
		}

		token, err := utils.GenerateJWT(admin.Email)
		if err != nil {
			return "", err
		}

		return token, nil

	} else {
		user, err := repositories.GetUserByEmail(email)

		if err != nil || !utils.CheckPasswordHash(password, user.Password) {
			return "", errors.New("invalid credentials")
		}
	
		if user.IsBlocked {
			return "", errors.New("user is blocked")
		}
	
		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			return "", err
		}
	
		return token, nil
	}
}


// GetEmailFromToken validates the token and returns the user's email
func GetEmailFromToken(token string) (string, error) {
    claims, err := utils.ValidateToken(token)
    if err != nil {
        return "", errors.New("invalid or expired token")
    }
    return claims.Email, nil
}

func BlockUser(email string) error {

	err := repositories.BlockUser(email)

	if err != nil {
		return errors.New("failed to block user")
	}

	return nil
}

func UnblockUser(email string) error {
	err := repositories.UnblockUser(email)
	if err != nil {
		return errors.New("failed to unblock user")
	}
	return nil
}

func GetUsersStatus() ([]models.UserResponse, error) {
	users, err := repositories.GetUsersStatus()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to get users status")
	}
	return users, nil
}