package http

import (
	"net/http"

	app "rest-go-service/restGoService/internal/app"

	"github.com/labstack/echo/v4"
)

type handlers struct {
	Logic app.LgPersons
}

func NewHandlers(e *echo.Echo, logics app.LgPersons) {

	h := &handlers{
		Logic: logics,
	}

	e.GET("/", home)
	e.POST("/person", h.addPerson)
	e.GET("/persons", h.getPersons)
	e.GET("/person/:id", h.getPerson)
	e.PUT("/person/:id", h.updatePerson)
	e.DELETE("/person/:id", h.deletePerson)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

func (l handlers) addPerson(c echo.Context) (err error) {
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	err = l.Logic.AddPerson(u)
	if err != nil {
		return err
	}

	return c.String(201, "New Person added")

}

func (l handlers) getPersons(c echo.Context) (err error) {

	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	result, err := l.Logic.GetPersons(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (l handlers) getPerson(c echo.Context) (err error) {
	id := c.Param("id")
	result, err := l.Logic.GetPerson(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (l handlers) updatePerson(c echo.Context) (err error) {
	id := c.Param("id")
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	err = l.Logic.UpdatePerson(u, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Data updated successfully")

}

func (l handlers) deletePerson(c echo.Context) (err error) {
	id := c.Param("id")
	err = l.Logic.DeletePerson(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Deleted Person with id - "+id)

}
