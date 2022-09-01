package main

import (
	"fmt"
	app "rest-go-service/restGoService/internal/app"
	http "rest-go-service/restGoService/internal/http"
	postgreSQL "rest-go-service/restGoService/internal/postgre"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// var personLogic app.LgPersons
var personDb app.DbPersons

func main() {

	e := echo.New()

	// lg := logic.NewLogic(personLogic)
	db := postgreSQL.NewDb(postgreSQL.NewConfig(), &logrus.Logger{}, personDb)

	http.NewHandlers(e, db)
	// //Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))

	fmt.Println(db)
}
