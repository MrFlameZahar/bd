package services

import (
	"fmt"

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
