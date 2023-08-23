package main

import (
	"errors"
	"github.com/Shemistan/Lesson_5/internal/api"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/service"
	"github.com/Shemistan/Lesson_5/internal/storage"
)

func main() {
	// Читай конфиг файл и вытаскивам различные данные
	conn := storage.NewConnect()
	db := storage.New("localhost", 5432, 30, conn)

	serv := service.New(db)
	handlres := api.New(serv)

	Server(handlres)
}

func Server(handlres api.IApi) {
	handlres.Add(&models.AddRequest{
		AuthParams: models.UserAuthParams{
			Login:    "login_1",
			Password: "password_1",
		},
		Date: models.UserDate{
			Name: "test_name",
			Age:  18,
		},
	})
}

func Sum(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
