package main

import (
	"fmt"
	"net/http"

	ports "main.go/internal/ports/http"
)

// AuthService
// - Регистрация
// - Авторизация
// - Восстановление пароля
// - подстверждение почты
// - права доступа

func main() {
	fmt.Println("Hello, World!")

	http.ListenAndServe(":8050", ports.NewRouter())
}
