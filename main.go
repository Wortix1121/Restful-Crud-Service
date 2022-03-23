package main

import (
	http "rest-go-service/restGoService/internal/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	http.Starting(e)

	//Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))
}
