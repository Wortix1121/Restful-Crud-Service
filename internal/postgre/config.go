package postgreSQL

//Config
type Config struct {
	Connect string
	Host    string
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		Connect: "host=localhost port=5432 user=postgres password=qwerty dbname=testdb sslmode=disable",
	}

}
