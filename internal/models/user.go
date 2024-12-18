package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"isVerified"`
}

type UserSignUp struct {
	Email    string
	Password string
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
