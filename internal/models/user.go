package models

import "time"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserSignUp struct {
	Email    string
	Password string
}

type UserVerification struct {
	IsVerified       bool
	VerificationCode string
	CodeExpireTime   time.Time
	Email            string
	ID               int
}
type UserVerificationRequest struct {
	VerificationCode string
	Email            string
}

func NewUserSignUp() *UserSignUp {
	return &UserSignUp{
		Email:    "",
		Password: "",
	}
}
func NewUser() *User {
	return &User{
		ID:    0,
		Email: "",
	}
}
