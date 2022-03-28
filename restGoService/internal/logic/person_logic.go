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

func (l logics) Add(u *app.Person, d app.DbPersons) (err error) {
	err = d.AddPerson(u)
	if err != nil {
		return err
	}
	return nil
}

func (l logics) Get(u *app.Person, p app.DbPersons) (all []app.Person, err error) {

	all, err = p.GetPersons(u)
	if err != nil {
		return all, nil
	}
	return all, err
}

func (l logics) GetAll(id string, p app.DbPersons) (getOne []app.Person, err error) {
	getOne, err = p.GetPerson(id)
	if err != nil {
		return getOne, nil
	}
	return getOne, err

}

func (l logics) Update(u *app.Person, id string, p app.DbPersons) (err error) {
	err = p.UpdatePerson(u, id)
	if err != nil {
		return err
	}
	return nil
}

func (l logics) Delete(id string, p app.DbPersons) (err error) {
	err = p.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}
