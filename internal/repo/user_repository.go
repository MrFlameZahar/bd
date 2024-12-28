package repo

import models "main.go/internal/models"

type UserRepository interface {
	FindUserByEmail(email string) (bool, error)
	AddUserToDB(userSignUp *models.UserSignUp) error
	GetUserFromDB(userSignUp *models.UserSignUp) (*models.User, error)
	GetUserAuthorisationFromDB(email string) (*models.UserSignUp, error)
}

type VerificationRepository interface {
	AddUserToVerificationDB(user *models.User, verificationCode string) error
	GetUserVerification(email string) (*models.UserVerification, error)
	ChangeVerificationState(userVerification *models.UserVerification, verificationState bool) error
	AddVerificationCode(userVerification *models.UserVerification, verificationCode string) error
}
