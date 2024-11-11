// services/auth_service.go
package services

import (
	"errors"
	"fmt"
	"auth-microservice/models"
	"auth-microservice/repositories"
	"auth-microservice/utils"
	"time"
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

func GeneratePasswordResetToken(email string) error {
	if _, err := repositories.GetUserByEmail(email); err != nil {
		return errors.New("user not found")
	}

	token, err := utils.GenerateJWT(email)
	if err != nil {
		return err
	}

	repositories.DeletePasswordResetToken(email)

	err = repositories.SavePasswordResetToken(email, token)
	if err != nil {
		return err
	}

	err = utils.SendPasswordResetEmail(email, token)
	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(email, password, token string) error {
	_, err := repositories.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}
	fmt.Println("Email: ", email)

	reset, err := repositories.GetPasswordResetToken(email)

	fmt.Println("Reset Password Token: ", reset.Token)
	fmt.Println("Created At: ", reset.CreatedAt)

	createdAt, err := time.Parse(time.RFC3339, reset.CreatedAt)
	if err != nil {
		fmt.Println("Error converting created_at to time.Time:", err)
		return err
	}
	

	if err != nil {
		return errors.New("invalid or expired token error")
	}

	fmt.Println("Token: ", token)
	fmt.Println("Reset Token: ", reset)
	fmt.Println("Password: ", password)

	if reset.Token != token {
		return errors.New("invalid or expired token diff token")
	}

	if createdAt.Add(15 * 60 * time.Second).Before(time.Now()) {
		return errors.New("expired token")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	err = repositories.UpdatePassword(email, hashedPassword)
	if err != nil {
		return err
	}

	repositories.DeletePasswordResetToken(email)

	return nil
}

func GeneratePin(email string) error {
	fmt.Println("Email: ", email)
	pin := utils.GenerateRandomString(6)
	fmt.Println("Pin: ", pin)
	repositories.DeletePin(email)
	err := repositories.SavePin(email, pin)

	if err != nil {
		return err
	}

	err = utils.SendPinEmail(email, pin)
	if err != nil {
		return err
	}

	return nil;
}

func VerifyPin(email, pin string) error {
	savedPin, err := repositories.GetPin(email)
	if err != nil {
		return errors.New("invalid pin")
	}

	if savedPin.Pin != pin {
		return errors.New("invalid pin")
	}

	repositories.DeletePin(email)

	createdAt, err := time.Parse(time.RFC3339, savedPin.CreatedAt)
	if err != nil {
		fmt.Println("Error converting created_at to time.Time:", err)
		return err
	}

	if createdAt.Add(1 * 60 * time.Second).Before(time.Now()) {
		return errors.New("expired pin")
	}

	return nil
}

func LoginUserWithGoogle(email string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
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