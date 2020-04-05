package dal

import (
	"fmt"

	"github.com/go-pg/pg/v9"
)

// DB is the global DB connection
var DB *pg.DB

// InitDb sets up the DB connection and saves a pointer to it in DB variable
func InitDb() {
	// DB
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "EveryTh1ngis4wesome",
		Database: "muni",
	})

	_, err := DB.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
	}
}
