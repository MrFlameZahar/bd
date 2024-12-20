package models

type RegisterResponse struct {
	JwtToken         string
	VerificationCode string
}
