package main

import (
	app "rest-go-service/restGoService/internal/app"
	http "rest-go-service/restGoService/internal/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var handler app.LgPersons
	http.NewHandlers(e, handler)
	//Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))
}
