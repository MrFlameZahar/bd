package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	models "main.go/internal/models"
	repo "main.go/internal/repo"
)

func SignUp(userSignUp *models.UserSignUp) (*models.User, error) {
	exist, err := repo.FindUserByEmail(userSignUp.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return &models.User{}, fmt.Errorf("пользователь уже существует")
	} else {
		err = repo.AddUserToDB(userSignUp)
		if err != nil {
			return nil, err
		}
	}

	user, err := repo.GetUserFromDB(userSignUp)

	if err != nil {
		return nil, err
	}

	err = repo.AddUserToVerificationDB(user, GenerateVerificationCode(user))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func LogIn(userSignUp *models.UserSignUp) (*models.User, error) {
	exist, err := repo.FindUserByEmail(userSignUp.Email)

	if err != nil {
		return nil, err
	}

	if exist {
		loginUser, err := repo.GetUserAuthorisationFromDB(userSignUp.Email)
		if err != nil {
			return nil, err
		}
		if loginUser.Email == userSignUp.Email && loginUser.Password == userSignUp.Password {
			user, err := repo.GetUserFromDB(userSignUp)

			if err != nil {
				return nil, err
			}
			return user, nil
		} else {
			return &models.User{}, fmt.Errorf("неправильный логин или пароль")
		}
	} else {
		return &models.User{}, fmt.Errorf("пользователь не зарегестрирован")
	}
}

func VerificationStatus(user *models.User) bool {
	verif, _ := repo.GetUserVerification(user)
	return verif.IsVerified
}

func Verify(userVerification *models.UserVerification, verificationCode string) error {
	valid, err := checkValidVerification(userVerification)
	if valid && err != nil {
		if userVerification.VerificationCode == verificationCode {
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
	if !userVerification.IsVerified {
		return false, fmt.Errorf("верификация не прошла проверку, почта уже верифицирована")
	}
	if userVerification.CodeExpireTime.Unix() > time.Now().Unix() {
		return true, nil
	} else {
		return false, fmt.Errorf("верификация не прошла проверку, истёк срок жизни кода верификации")
	}
}

func GenerateVerificationCode(user *models.User) string {
	hash := md5.New()
	hash.Write([]byte(user.Email))
	hashBytes := hash.Sum(nil)

	// Преобразуем хеш в строку
	return hex.EncodeToString(hashBytes)
}
