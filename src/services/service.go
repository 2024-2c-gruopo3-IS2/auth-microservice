// services/auth_service.go
package services

import (
	"errors"

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
	user, err := repositories.GetUserByEmail(email, isAdmin)
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


// GetEmailFromToken validates the token and returns the user's email
func GetEmailFromToken(token string) (string, error) {
    claims, err := utils.ValidateToken(token)
    if err != nil {
        return "", errors.New("invalid or expired token")
    }
    return claims.Email, nil
}

func BlockUser(email string) error {
	query := `UPDATE users SET is_blocked = TRUE WHERE email = $1`
	result, err := config.DB.Exec(query, email)
	if err != nil {
		return fmt.Errorf("failed to block user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not determine affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
