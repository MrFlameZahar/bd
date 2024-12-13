package models

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
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
