package dal

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

// DB is the global DB connection
var DB *pg.DB

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// InitDb sets up the DB connection and saves a pointer to it in DB variable
func InitDb() {
	// DB
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "EveryTh1ngis4wesome",
		Database: "muni",
	})
	DB.AddQueryHook(dbLogger{})

	_, err := DB.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
	}
}
