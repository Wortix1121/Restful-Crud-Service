package main

import (
	app "rest-go-service/restGoService/internal/app"
	http "rest-go-service/restGoService/internal/http"

	"github.com/labstack/echo/v4"
)

var handler app.LgPersons

func main() {
	e := echo.New()

	http.NewHandlers(e, handler)
	//Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))
}
