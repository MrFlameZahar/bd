package services

import (
	models "main.go/internal/models"
	repo "main.go/internal/repo"
)

func SignUp(userSignUp *models.UserSignUp) (*models.User, error) {

	user, err := repo.CreateUser(userSignUp)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func LogIn(userSignUp *models.UserSignUp) (*models.User, error) {

	user, err := repo.LoginUser(userSignUp)

	if err != nil {
		return nil, err
	}

	return user, nil

}
