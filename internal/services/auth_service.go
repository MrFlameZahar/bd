package services

import (
	"fmt"

	models "main.go/internal/models"
	"main.go/internal/repo"
)

type UserData struct {
	repo repo.UserRepository
}

func NewUserData(repo repo.UserRepository) *UserData {
	return &UserData{repo: repo}
}

func SignUp(userSignUp *models.UserSignUp) (*models.User, error) {
	userData := NewUserData(repo.NewUserDataBase())

	exist, err := userData.repo.FindUserByEmail(userSignUp.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return &models.User{}, fmt.Errorf("пользователь уже существует")
	} else {
		err = userData.repo.AddUserToDB(userSignUp)
		if err != nil {
			return nil, err
		}
	}

	user, err := userData.repo.GetUserFromDB(userSignUp)

	if err != nil {
		return nil, err
	}

	AddUserVerificationData(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func LogIn(userSignUp *models.UserSignUp) (*models.User, error) {
	userData := NewUserData(repo.NewUserDataBase())

	exist, err := userData.repo.FindUserByEmail(userSignUp.Email)

	if err != nil {
		return nil, err
	}

	if exist {
		loginUser, err := userData.repo.GetUserAuthorisationFromDB(userSignUp.Email)
		if err != nil {
			return nil, err
		}
		if loginUser.Email == userSignUp.Email && loginUser.Password == userSignUp.Password {
			if VerificationStatus(userSignUp.Email) {
				user, err := userData.repo.GetUserFromDB(userSignUp)

				if err != nil {
					return nil, err
				}
				return user, nil
			} else {
				return &models.User{}, fmt.Errorf("почта не верифицирована")
			}
		} else {
			return &models.User{}, fmt.Errorf("неправильный логин или пароль")
		}
	} else {
		return &models.User{}, fmt.Errorf("пользователь не зарегестрирован")
	}
}
