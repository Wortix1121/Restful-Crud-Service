package http

import (
	"net/http"

	app "rest-go-service/restGoService/internal/app"

	"github.com/labstack/echo/v4"
)

type handlers struct {
	Handler app.LgPersons
}

func NewHandlers(e *echo.Echo, handler app.LgPersons) {

	h := &handlers{
		Handler: handler,
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

func (h handlers) addPerson(c echo.Context) (err error) {
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	err = h.Handler.AddPerson(u)
	if err != nil {
		return err
	}

	return c.String(201, "New Person added")

}

func (h handlers) getPersons(c echo.Context) (err error) {

	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}

	result, err := h.Handler.GetPersons(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h handlers) getPerson(c echo.Context) (err error) {
	id := c.Param("id")
	result, err := h.Handler.GetPerson(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h handlers) updatePerson(c echo.Context) (err error) {
	id := c.Param("id")
	u := new(app.Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	err = h.Handler.UpdatePerson(u, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Data updated successfully")

}

func (h handlers) deletePerson(c echo.Context) (err error) {
	id := c.Param("id")
	err = h.Handler.DeletePerson(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Deleted Person with id - "+id)

}
