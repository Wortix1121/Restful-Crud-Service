package http

import (
	"net/http"

	. "rest-go-service/restGoService/internal/app"
	logic "rest-go-service/restGoService/internal/logic"

	"github.com/labstack/echo/v4"
)

func Starting() {
	e := echo.New()

	e.GET("/", home)
	e.POST("/person", addPerson)
	e.GET("/persons", getPersons)
	e.GET("/person/:id", getPerson)
	e.PUT("/person/:id", updatePerson)
	e.DELETE("/person/:id", deletePerson)

	//Запуск локального сервера
	e.Logger.Fatal(e.Start(":8000"))
}
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

//Метод -POST-
//Добавление нового Person
func addPerson(c echo.Context) error {
	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	logic.AddPerson(u)

	return c.String(201, "New Person added")

}

func getPersons(c echo.Context) (err error) {
	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	result, err := logic.GetPersons(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func getPerson(c echo.Context) (err error) {
	id := c.Param("id")
	result, err := logic.GetPerson(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func updatePerson(c echo.Context) error {
	id := c.Param("id")
	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	logic.UpdatePerson(u, id)

	return c.String(http.StatusOK, "Data updated successfully")

}

func deletePerson(c echo.Context) error {
	id := c.Param("id")

	logic.DeletePerson(id)
	return c.String(http.StatusOK, "Deleted Person with id - "+id)

}
