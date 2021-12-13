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

func Connect() *sql.DB {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")

	return db
}

func AddPerson(u *Person) {
	//Create Person
	db := Connect()
	sqlInq := "INSERT INTO persons(email, phone, firstName, lastName)VALUES($1,$2,$3,$4)"
	create, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName)
	if err != nil {
		log.Fatal("Error creating Person", err.Error())
	}
	defer create.Close()
}

func GetPersons(u *Person) JsonPerson {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstName, lastName FROM persons ORDER BY id asc"
	get, err := db.Query(sqlInq)
	if err != nil {
		log.Fatal("Error creating Person", err.Error())
	}
	defer get.Close()

	result := JsonPerson{}

	for get.Next() {
		persons := Person{}
		err2 := get.Scan(&persons.Id, &persons.Email, &persons.Phone, &persons.FirstName, &persons.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		result.JsonPersons = append(result.JsonPersons, persons)
	}
	return result

}

func GetPerson(id string) JsonPerson {
	db := Connect()
	sqlInq := "SELECT id, email, phone, firstname, lastname FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	if err != nil {
		log.Println("Error", err.Error())
	}
	defer res.Close()

	result := JsonPerson{}

	for res.Next() {
		persons := Person{}
		err2 := res.Scan(&persons.Id, &persons.Email, &persons.Phone, &persons.FirstName, &persons.LastName)
		if err2 != nil {
			log.Println(err2)
		}
		result.JsonPersons = append(result.JsonPersons, persons)
	}
	return result
}

func UpdatePerson(u *Person, id string) {
	db := Connect()
	sqlInq := "UPDATE persons SET email=$1, phone=$2, firstname=$3, lastname=$4 WHERE id=$5"
	res, err := db.Query(sqlInq, u.Email, u.Phone, u.FirstName, u.LastName, id)
	if err != nil {
		log.Println(err)
	}
	defer res.Close()
}

func DeletePerson(id string) {
	db := Connect()
	sqlInq := "DELETE FROM persons WHERE id=$1"
	res, err := db.Query(sqlInq, id)
	if err != nil {
		log.Println(err)
	}
	defer res.Close()
}
