package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	models "main.go/internal/models"
	repo "main.go/internal/repo"
)

// internal "main.go/internal"
func VerificationStatus(email string) bool {
	user, _ := repo.GetUserVerification(email)
	return user.IsVerified
}

func Verify(userVerification *models.UserVerification) error {
	valid, err := checkValidVerification(userVerification)
	if valid && err != nil {
		if userVerification.VerificationCode == GenerateVerificationCode(userVerification.Email) {
			repo.ChangeVerificationState(userVerification, true)
			return err
		} else {
			return err
		}
	} else {
		return err
	}
}

func checkValidVerification(userVerification *models.UserVerification) (bool, error) {
	if userVerification.IsVerified {
		return false, fmt.Errorf("верификация не прошла проверку, почта уже верифицирована")
	}
	if userVerification.CodeExpireTime.Unix() > time.Now().Unix() {
		return true, nil
	} else {
		return false, fmt.Errorf("верификация не прошла проверку, истёк срок жизни кода верификации")
	}
}

func GenerateVerificationCode(email string) string {
	hash := md5.New()
	hash.Write([]byte(email))
	hashBytes := hash.Sum(nil)

	// Преобразуем хеш в строку
	return hex.EncodeToString(hashBytes)
}
