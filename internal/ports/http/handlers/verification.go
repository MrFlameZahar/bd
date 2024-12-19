package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Verification(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("hello world")
	fmt.Fprintf(w, "Добро пожаловать авторизированный пользователь")
}
