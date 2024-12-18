package repo

import (
	"database/sql"

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

// Перенести соединение в общее
