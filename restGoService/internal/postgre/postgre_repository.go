package postgreSQL

import (
	"database/sql"
	"fmt"
	"log"

	app "rest-go-service/restGoService/internal/app"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// type Server interface {
// }

// type logics struct {
// 	person app.Persons
// }

// func NewLogic(person app.Persons) *logics {
// 	return &logics{person: person}
// }

type dataBase struct {
	config *Config
	logger *logrus.Logger
	person app.DbPersons
}

func NewDb(config *Config, logger *logrus.Logger, person app.DbPersons) *dataBase {
	return &dataBase{
		config: config,
		logger: logger,
		person: person,
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty"
	dbname   = "testdb"
)

var (
	Persons = app.Person{}
)

func Connect() *sql.DB {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")

	return db
}

func (d dataBase) AddPerson(u *app.Person) (err error) {
	//Create Person
	db := Connect()
	sqlInq := "INSERT INTO persons(email, phone, firstName, lastName)VALUES($1,$2,$3,$4)"
	create, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName)
	if err != nil {
		return err
	}
	defer create.Close()
	return nil

}

func (d dataBase) GetPersons(u *app.Person) (persons []app.Person, err error) {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstName, lastName FROM persons ORDER BY id asc"
	get, err := db.Query(sqlInq)
	if err != nil {
		log.Fatal("Error creating Person", err.Error())
	}
	defer get.Close()

	persons = []app.Person{}

	for get.Next() {
		var person app.Person
		err2 := get.Scan(&person.Id, &person.Email, &person.Phone, &person.FirstName, &person.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		persons = append(persons, person)
	}
	return persons, nil

}

func (d dataBase) GetPerson(id string) (persons app.Person, err error) {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstname, lastname FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	if err != nil {
		log.Println("Error", err.Error())
	}
	defer res.Close()

	Persons = app.Person{}

	for res.Next() {
		var person app.Person
		err2 := res.Scan(&person.Id, &person.Email, &person.Phone, &person.FirstName, &person.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		Persons = person
	}
	return Persons, nil
}

func (d dataBase) UpdatePerson(u *app.Person, id string) (err error) {
	db := Connect()
	sqlInq := "UPDATE persons SET email=$1, phone=$2, firstname=$3, lastname=$4 WHERE id=$5"
	res, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName, id)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}

func (d dataBase) DeletePerson(id string) (err error) {
	db := Connect()
	sqlInq := "DELETE FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}
