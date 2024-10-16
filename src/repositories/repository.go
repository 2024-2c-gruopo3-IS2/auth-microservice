// repositories/auth_repository.go
package repositories

import (
	"errors"

	"auth-microservice/config"
	"auth-microservice/models"
)

func CreateUser(user *models.User, isAdmin bool) error {
	table := getTable(isAdmin)
	query := "INSERT INTO " + table + " (email, password) VALUES ($1, $2)"
	_, err := config.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		return errors.New("User already exists")
	}
	return nil
}

func GetAdminByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT email, password FROM admins WHERE email=$1"
	err := config.DB.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT email, password, is_blocked FROM users WHERE email=$1"
	err := config.DB.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func BlockUser(email string) error {

	query := `UPDATE users SET is_blocked = TRUE WHERE email = $1`
	_, err := config.DB.Exec(query, email)
	if err != nil {
		return errors.New("failed to block user")
	}
	return nil
}

func UnblockUser(email string) error {
	query := `UPDATE users SET is_blocked = FALSE WHERE email = $1`
	_, err := config.DB.Exec(query, email)
	if err != nil {
		return errors.New("failed to unblock user")
	}
	return nil
}

func GetUsersStatus() ([]models.UserResponse, error) {
	var users []models.UserResponse
	query := `SELECT email, is_blocked FROM users`
	err := config.DB.Select(&users, query)
	if err != nil {
		return nil, errors.New("failed to get users status db")
	}
	return users, nil
}

func getTable(isAdmin bool) string {
	if isAdmin {
		return "admins"
	}
	return "users"
}
