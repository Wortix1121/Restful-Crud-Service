package app

type Person struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LgPersons interface {
	AddPerson(u *Person) error
	GetPersons(u *Person) ([]Person, error)
	GetPerson(id string) (Person, error)
	UpdatePerson(u *Person, id string) error
	DeletePerson(id string) error
}

type DbPersons interface {
	AddPerson(u *Person) error
	GetPersons(u *Person) ([]Person, error)
	GetPerson(id string) (Person, error)
	UpdatePerson(u *Person, id string) error
	DeletePerson(id string) error
}

// type JsonPerson struct {
// 	JsonPersons []Person `json:"jsonpersons"`
// }
