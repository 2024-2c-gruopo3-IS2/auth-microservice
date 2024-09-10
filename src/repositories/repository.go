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
		return errors.New("error creating user: " + err.Error())
	}
	return nil
}

func GetUserByEmail(email string, isAdmin bool) (*models.User, error) {
	var user models.User
	table := getTable(isAdmin)
	query := "SELECT email, password FROM " + table + " WHERE email=$1"
	err := config.DB.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func getTable(isAdmin bool) string {
	if isAdmin {
		return "admins"
	}
	return "users"
}