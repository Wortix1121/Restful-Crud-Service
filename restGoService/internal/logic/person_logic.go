package logic

import (
	app "rest-go-service/restGoService/internal/app"
)

type logics struct {
	person app.LgPersons
}

func NewLogic(person app.LgPersons) *logics {
	return &logics{person: person}
}

func (l logics) AddPerson(u *app.Person, p app.DbPersons) (err error) {
	err = p.AddPerson(u)
	if err != nil {
		return err
	}
	return nil
}

func (l logics) GetPerson(id string, p app.DbPersons) (getOne app.Person, err error) {

	getOne, err = p.GetPerson(id)
	if err != nil {
		return getOne, err
	}
	return getOne, nil
}

func (l logics) GetPersons(u *app.Person, p app.DbPersons) (all []app.Person, err error) {
	all, err = p.GetPersons(u)
	if err != nil {
		return all, err
	}
	return all, nil

}

func (l logics) UpdatePerson(u *app.Person, id string, p app.DbPersons) (err error) {
	err = p.UpdatePerson(u, id)
	if err != nil {
		return err
	}
	return nil
}

func (l logics) DeletePerson(id string, p app.DbPersons) (err error) {
	err = p.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}
