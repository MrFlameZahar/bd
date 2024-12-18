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

	}
	// FindUserByEmail(userSignUp.Email)
	// If пользователя нет, то создаем пользователя, отправляем запрос в repo на создание его в базе, если ошибок нет, то возвращаем пользователя
	// Else возвращаем ошибку что пользователь уже есть и пустую модель юзера

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
