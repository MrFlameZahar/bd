package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/mail"

	models "main.go/internal/models"
	services "main.go/internal/services"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	// Чтение и парсинг тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var userSignUp models.UserSignUp
	// Декодирование JSON
	err = json.Unmarshal(body, &userSignUp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidEmail(userSignUp.Email) {
		return
	}

	user, err := services.SignUp(&userSignUp)

	if err != nil {
		fmt.Println("Ошибка регистрации:", err)
		fmt.Fprintf(w, "Ошибка регистрации: %v", err)
		return
	}

	token, err := services.GenerateJWT(user, services.VerificationStatus(user))
	if err != nil {

		fmt.Println("Error generating token:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Код ответа
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Вы авторизованы на сайте")
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
