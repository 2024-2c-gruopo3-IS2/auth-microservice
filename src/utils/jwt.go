// utils/jwt.go
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define una clave secreta para firmar los JWT
var jwtKey = []byte("snap_msg_secret_key")

// Claims define la estructura de los claims del token
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT genera un token JWT para un usuario
func GenerateJWT(userEmail string) (string, error) {
	// Configura los claims del JWT
	claims := &Claims{
		Email: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Crea el token con los claims especificados
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken valida un JWT y devuelve los claims si es válido
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}