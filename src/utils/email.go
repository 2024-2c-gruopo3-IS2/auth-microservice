package utils

import (
	"fmt"
	"os"
	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(email string, token string) error {
	body := fmt.Sprintf("<p>Para restablecer tu contraseña, usa el siguiente token:</p><p><strong>%s</strong></p>", token)

	smtpUser := os.Getenv("SENDING_EMAIL")
	fmt.Println(smtpUser)
    smtpPass := os.Getenv("SENDING_EMAIL_PASSWORD")
	fmt.Println(smtpPass)

	mail := gomail.NewMessage()
	mail.SetHeader("From", smtpUser)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Password Reset Request")
	mail.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, smtpUser, smtpPass)

	print("Dialer: ", dialer)

	err := dialer.DialAndSend(mail)

	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}