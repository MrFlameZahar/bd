package ports

import (
	"net/http"

	"github.com/gorilla/mux"
	handlers "main.go/internal/ports/http/handlers"
	"main.go/internal/services"
)

// Инициализация роутера

func NewRouter() http.Handler {

	authMux := mux.NewRouter()
	authMux.HandleFunc("/login", handlers.Login)
	authMux.HandleFunc("/register", handlers.SignUp).Methods("POST")
	authMux.Handle("/", services.AuthMiddleware(http.HandlerFunc(handlers.MainPage)))

	return authMux
}
