package postgreSQL

import (
	"database/sql"
	"log"

	"fmt"

	. "rest-go-service/restGoService/internal/app"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty"
	dbname   = "testdb"
)

var (
	onePerson = []Person{}
	persons   = []Person{}
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

func AddPerson(u *Person) (err error) {
	//Create Person
	db := Connect()
	sqlInq := "INSERT INTO persons(email, phone, firstName, lastName)VALUES($1,$2,$3,$4)"
	create, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName)
	defer create.Close()
	return err
}

func GetPersons(u *Person) (persons []Person, err error) {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstName, lastName FROM persons ORDER BY id asc"
	get, err := db.Query(sqlInq)
	if err != nil {
		log.Fatal("Error creating Person", err.Error())
	}
	defer get.Close()

	persons = []Person{}

	for get.Next() {
		var person Person
		err2 := get.Scan(&person.Id, &person.Email, &person.Phone, &person.FirstName, &person.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		persons = append(persons, person)
	}
	return persons, err

}

func GetPerson(id string) (persons []Person, err error) {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstname, lastname FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	if err != nil {
		log.Println("Error", err.Error())
	}
	defer res.Close()

	onePerson = []Person{}

	for res.Next() {
		var persons Person
		err2 := res.Scan(&persons.Id, &persons.Email, &persons.Phone, &persons.FirstName, &persons.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		onePerson = append(onePerson, persons)
	}
	return onePerson, err
}

func UpdatePerson(u *Person, id string) (err error) {
	db := Connect()
	sqlInq := "UPDATE persons SET email=$1, phone=$2, firstname=$3, lastname=$4 WHERE id=$5"
	res, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName, id)
	defer res.Close()
	return err
}

func DeletePerson(id string) (err error) {
	db := Connect()
	sqlInq := "DELETE FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	defer res.Close()
	return err
}
