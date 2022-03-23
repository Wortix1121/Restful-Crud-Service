package http

import (
	"net/http"

	app "rest-go-service/restGoService/internal/app"

	"github.com/labstack/echo/v4"
)

func Starting(e *echo.Echo) {

	e.GET("/", home)
	e.POST("/person", addPerson)
	e.GET("/persons", getPersons)
	e.GET("/person/:id", getPerson)
	e.PUT("/person/:id", updatePerson)
	e.DELETE("/person/:id", deletePerson)

}
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

func addPerson(c echo.Context, l app.LgPersons) (err error) {
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	err = l.Add(u)
	if err != nil {
		return err
	}

	return c.String(201, "New Person added")

}

func getPersons(c echo.Context, l app.LgPersons) (err error) {

	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	result, err := l.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func getPerson(c echo.Context, l app.LgPersons) (err error) {
	id := c.Param("id")
	result, err := l.GetAll(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func updatePerson(c echo.Context, l app.LgPersons) (err error) {
	id := c.Param("id")
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	err = l.Update(u, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Data updated successfully")

}

func deletePerson(c echo.Context, l app.LgPersons) (err error) {
	id := c.Param("id")
	err = l.Delete(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Deleted Person with id - "+id)

}
