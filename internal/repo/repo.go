package repo

import (
	"database/sql"
	"fmt"

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
func CreateUser(userSignUp *models.UserSignUp) (*models.User, error) {

	user := models.NewUser()

	_, err := DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", userSignUp.Email, userSignUp.Password)
	DB.QueryRow("SELECT email, id FROM users WHERE email = $1", userSignUp.Email).Scan(&user.Email, &user.ID)

	if err != nil {
		return &models.User{}, err
	}
	return user, nil

}

func LoginUser(userSignUp *models.UserSignUp) (*models.User, error) {
	exist, _ := FindUserByEmail(userSignUp.Email)
	// Избавиться от условной логики здесь и унести её в ауф сервайс
	if exist {
		var password string
		var email string
		DB.QueryRow("SELECT email, password FROM users WHERE email = $1", userSignUp.Email).Scan(&email, &password)
		if userSignUp.Email == email && userSignUp.Password == password {
			user := models.NewUser()
			DB.QueryRow("SELECT email, id FROM users WHERE email = $1", userSignUp.Email).Scan(&user.Email, &user.ID)
			return user, nil
		} else {
			return &models.User{}, fmt.Errorf("неправильный логин или пароль")
		}

	} else {
		return &models.User{}, fmt.Errorf("пользователь не существует")
	}
}

func UpdateUser(db *sql.DB) {

	db.Query("UPDATE users SET name = 'new_name' WHERE id = 1")
}

// Перенести соединение в общее
