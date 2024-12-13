package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	secretKey "main.go/internal"
	models "main.go/internal/models"
)

func GenerateJWT(user *models.User) (string, error) {
	// Создание claims (данных для токена)
	email := user.Email
	id := user.ID

	claims := jwt.MapClaims{
		"sub":  id,                               // Subject (например, user ID)
		"name": email,                            // Имя пользователя
		"iat":  time.Now().Unix(),                // Время создания токена
		"exp":  time.Now().Add(time.Hour).Unix(), // Время истечения токена
	}

	// Создание токена с данными и алгоритмом подписи
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подпись токена секретным ключом
	tokenString, err := token.SignedString(secretKey.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
