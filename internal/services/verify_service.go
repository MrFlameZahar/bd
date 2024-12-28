package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/go-gomail/gomail"
	"main.go/internal"
	models "main.go/internal/models"
	"main.go/internal/repo"
)

type VerificationData struct {
	repo repo.VerificationRepository
}

func NewVerificationData(repo repo.VerificationRepository) *VerificationData {
	return &VerificationData{repo: repo}
}

// internal "main.go/internal"
func VerificationStatus(email string) bool {
	verificationData := NewVerificationData(repo.VerificationDataBase{})

	user, _ := verificationData.repo.GetUserVerification(email)
	return user.IsVerified
}

func AddUserVerificationData(user *models.User) error {
	err := repo.AddUserToVerificationDB(user, GenerateVerificationCode(user.Email))
	return err
}

func Verify(userVerification *models.UserVerification) error {
	valid, err := checkValidVerification(userVerification)
	if valid && err == nil {
		if userVerification.VerificationCode == GenerateVerificationCode(userVerification.Email) {
			repo.ChangeVerificationState(userVerification, true)
			return nil
		} else {
			return fmt.Errorf("неправильный код верификации")
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

func SendVerificationLetter(email string, verificationCode string) {
	smtpHost := "smtp.yandex.ru"
	smtpPort := 465

	message := gomail.NewMessage()
	message.SetHeader("From", internal.YandexMailLogin)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Верификация почты")
	message.SetBody("text/plain", "Ваш код верификации: "+verificationCode)

	dialer := gomail.NewDialer(smtpHost, smtpPort, internal.YandexMailLogin, internal.YandexMailPassword)
	dialer.SSL = true

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("Не удалось отправить письмо: %v", err)
	}

	log.Println("Письмо успешно отправлено!")
}
