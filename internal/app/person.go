package app

type Person struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// type JsonPerson struct {
// 	JsonPersons []Person `json:"jsonpersons"`
// }
