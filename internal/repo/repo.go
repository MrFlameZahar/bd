package repo

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	models "main.go/internal/models"
)

var DB *sql.DB

func ConnectDB() error {
	connStr := "user=mrflame password=Zaxaro12 dbname=test host=127.0.0.1 port=5432 sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	return err
}

func FindUserByEmail(email string) (bool, error) {

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`

	err := DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Изучи -> pgx5
func AddUserToDB(userSignUp *models.UserSignUp) error {
	_, err := DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", userSignUp.Email, userSignUp.Password)

	if err != nil {
		return err
	}
	return nil
}

func GetUserFromDB(userSignUp *models.UserSignUp) (*models.User, error) {
	user := models.NewUser()

	err := DB.QueryRow("SELECT email, id FROM users WHERE email = $1", userSignUp.Email).Scan(&user.Email, &user.ID)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func GetUserAuthorisationFromDB(email string) (*models.UserSignUp, error) {
	userSignUp := models.NewUserSignUp()
	err := DB.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&userSignUp.Email, &userSignUp.Password)
	if err != nil {
		return &models.UserSignUp{}, err
	}
	return userSignUp, nil
}

func AddUserToVerificationDB(user *models.User, verificationCode string) error {
	_, err := DB.Exec("INSERT INTO user_verification (is_verified, verification_code, id, email, code_expire_time) VALUES ($1, $2, $3, $4, $5)", false, verificationCode, user.ID, user.Email, time.Now().Add(time.Hour))
	return err
}

func GetUserVerification(email string) (*models.UserVerification, error) {
	var userVerification models.UserVerification

	err := DB.QueryRow("SELECT is_verified, verification_code, id, email, code_expire_time FROM user_verification WHERE email = $1", email).Scan(&userVerification.IsVerified, &userVerification.VerificationCode, &userVerification.ID, &userVerification.Email, &userVerification.CodeExpireTime)

	if err != nil {
		return &userVerification, err
	}

	return &userVerification, nil
}

func ChangeVerificationState(userVerification *models.UserVerification, verificationState bool) error {
	_, err := DB.Exec("UPDATE user_verification SET is_verified = $1 WHERE ID = $2", verificationState, userVerification.ID)

	return err
}

func AddVerificationCode(userVerification *models.UserVerification, verificationCode string) error {
	time := time.Now().Add(time.Hour)

	_, err := DB.Exec("UPDATE user_verification SET verification_code = $1, code_expire_time = $2 WHERE ID = $3", verificationCode, time, userVerification.ID)

	return err
}
