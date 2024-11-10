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

func SavePasswordResetToken(email, token string) error {
	query := `INSERT INTO password_resets (email, token, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP)`
	_, err := config.DB.Exec(query, email, token)
	if err != nil {
		return errors.New("failed to save password reset token")
	}
	return nil
}

func GetPasswordResetToken(email string) (models.ResetResponse, error) {
	var reset models.ResetResponse
	query := `SELECT token, created_at FROM password_resets WHERE email = $1`
	err := config.DB.Get(&reset, query, email)
	if err != nil {
		return reset, errors.New("failed to get password reset token")
	}
	return reset, nil
}

func UpdatePassword(email, password string) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`
	_, err := config.DB.Exec(query, password, email)
	if err != nil {
		return errors.New("failed to update password")
	}
	return nil
}

func DeletePasswordResetToken(email string) error {
	query := `DELETE FROM password_resets WHERE email = $1`
	_, err := config.DB.Exec(query, email)
	if err != nil {
		return errors.New("failed to delete password reset token")
	}
	return nil
}

func SavePin(email, pin string) error {
	query := `INSERT INTO pins (email, pin, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP)`
	_, err := config.DB.Exec(query, email, pin)
	if err != nil {
		return errors.New("failed to save pin")
	}
	return nil
}

func GetPin(email string) (models.PinResponse, error) {
	var pin models.PinResponse
	query := `SELECT pin, created_at FROM pins WHERE email = $1`
	err := config.DB.Get(&pin, query, email)
	if err != nil {
		return pin, errors.New("failed to get pin")
	}
	return pin, nil
}

func DeletePin(email string) error {
	query := `DELETE FROM pins WHERE email = $1`
	_, err := config.DB.Exec(query, email)
	if err != nil {
		return errors.New("failed to delete pin")
	}
	return nil
}