package main

import (
	"fmt"
	"net/http"

	ports "main.go/internal/ports/http"
	repo "main.go/internal/repo"
)

func main() {
	repo.ConnectDB()

	fmt.Println("Hello, World!")

	http.ListenAndServe(":8050", ports.NewRouter())

}

//То ду

// - подтверждение почты
// - Создание вериф кода во время регистрации
// - отправка ссылки для подтверждения на почту
// - Настройка истечения актуальности кода верификации
// - Сурогатная ручка которая возвращает статус верификации
// - Ветка в гитхабе
// - Добавление поля isVerified в клеймы токена
