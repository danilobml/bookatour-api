package utils

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email, userId, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"role": role,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	return token.SignedString(jwtKey)
}
