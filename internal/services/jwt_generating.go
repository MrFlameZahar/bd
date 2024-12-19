package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	secretKey "main.go/internal"
	models "main.go/internal/models"
)

func GenerateJWT(user *models.User, verified bool) (string, error) {
	// Создание claims (данных для токена)
	email := user.Email
	id := user.ID
	claims := jwt.MapClaims{
		"sub":        id,
		"email":      email,
		"isVerified": verified,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
