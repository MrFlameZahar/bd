package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	models "main.go/internal/models"
	"main.go/internal/repo"
	services "main.go/internal/services"
)

func Verification(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var userVerificationRequest models.UserVerificationRequest

	err = json.Unmarshal(body, &userVerificationRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if services.VerificationStatus(userVerificationRequest.Email) {
		fmt.Fprintf(w, "Почта уже верифицирована")
		return
	} else {
		userVerification, err := repo.GetUserVerification(userVerificationRequest.Email)
		if err != nil {
			fmt.Fprintf(w, "Ошибка верификации: %v", err)
			return
		}

		err = services.Verify(userVerification)

		if err != nil {
			fmt.Fprintf(w, "Ошибка верификации: %v", err)
			return
		}
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode("hello world")
	fmt.Fprintf(w, "Верификация пройдена")
}
