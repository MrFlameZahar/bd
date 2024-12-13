package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	models "main.go/internal/models"
)

func FindUserByEmail(email string, db *sql.DB) (bool, error) {

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`

	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Изучи -> pgx5
func CreateUser(userSignUp *models.UserSignUp) (*models.User, error) {

	db, err := connectDB()

	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	defer db.Close()

	exist, _ := FindUserByEmail(userSignUp.Email, db)
	if exist {
		return &models.User{}, fmt.Errorf("пользователь уже существует")
	} else {
		user := models.NewUser()

		_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", userSignUp.Email, userSignUp.Password)
		db.QueryRow("SELECT email, id FROM users WHERE email = $1", userSignUp.Email).Scan(&user.Email, &user.ID)

		if err != nil {
			return &models.User{}, err
		}
		return user, nil
	}

}

func LoginUser(userSignUp *models.UserSignUp) (*models.User, error) {
	db, err := connectDB()

	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	exist, _ := FindUserByEmail(userSignUp.Email, db)

	if exist {
		var password string
		var email string
		db.QueryRow("SELECT email, password FROM users WHERE email = $1", userSignUp.Email).Scan(&email, &password)
		fmt.Println(email, password, userSignUp.Email, userSignUp.Password)
		if userSignUp.Email == email {
			if userSignUp.Password == password {
				user := models.NewUser()
				db.QueryRow("SELECT email, id FROM users WHERE email = $1", userSignUp.Email).Scan(&user.Email, &user.ID)
				return user, nil
			} else {
				return &models.User{}, fmt.Errorf("неправильный пароль")
			}
		} else {
			return &models.User{}, fmt.Errorf("неправильный логин")
		}

	} else {
		return &models.User{}, fmt.Errorf("пользователь не существует")
	}
}

func UpdateUser(db *sql.DB) {

	db.Query("UPDATE users SET name = 'new_name' WHERE id = 1")
}

func connectDB() (*sql.DB, error) {
	connStr := "user=mrflame password=Zaxaro12 dbname=test host=127.0.0.1 port=5432 sslmode=disable"

	// Открываем соединение
	return sql.Open("postgres", connStr)
}
