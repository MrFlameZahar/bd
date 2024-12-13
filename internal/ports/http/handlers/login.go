package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	models "main.go/internal/models"
	services "main.go/internal/services"
)

func Login(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	var userSignUp models.UserSignUp

	// Декодирование JSON
	err = json.Unmarshal(body, &userSignUp)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var user *models.User

	user, err = services.LogIn(&userSignUp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		// Код ответа
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Неправильный логин или пароль")

		fmt.Println("Error login:", err)
		return
	}
	token, err := services.GenerateJWT(user)
	if err != nil {

		fmt.Println("Error generating token:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Код ответа
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
