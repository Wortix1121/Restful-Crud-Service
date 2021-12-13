package logic

import (
	. "rest-go-service/restGoService/internal/app"
	"rest-go-service/restGoService/internal/postgreSQL"
)

func AddPerson(u *Person) {
	postgreSQL.AddPerson(u)

}

func GetPersons(u *Person) JsonPerson {
	all := postgreSQL.GetPersons(u)
	return all
}

func GetPerson(id string) JsonPerson {
	getOne := postgreSQL.GetPerson(id)
	return getOne

}

func UpdatePerson(u *Person, id string) {
	postgreSQL.UpdatePerson(u, id)

}

func DeletePerson(id string) {
	postgreSQL.DeletePerson(id)

}
