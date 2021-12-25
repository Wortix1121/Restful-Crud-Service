package logic

import (
	. "rest-go-service/restGoService/internal/app"
	postgreSQL "rest-go-service/restGoService/internal/postgre"
)

func AddPerson(u *Person) (err error) {
	err = postgreSQL.AddPerson(u)
	if err != nil {
		return err
	}
	return nil
}

func GetPersons(u *Person) (all []Person, err error) {
	all, err = postgreSQL.GetPersons(u)
	if err != nil {
		return all, nil
	}
	return all, err
}

func GetPerson(id string) (getOne []Person, err error) {
	getOne, err = postgreSQL.GetPerson(id)
	if err != nil {
		return getOne, nil
	}
	return getOne, err

}

func UpdatePerson(u *Person, id string) (err error) {
	err = postgreSQL.UpdatePerson(u, id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id string) (err error) {
	err = postgreSQL.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}
