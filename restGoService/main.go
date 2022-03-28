package main

import (
	app "rest-go-service/restGoService/internal/app"
	http "rest-go-service/restGoService/internal/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var logic app.LgPersons
	http.NewHandlers(e, logic)
	//Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))
}
