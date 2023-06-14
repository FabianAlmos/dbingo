package db

import (
	"database/sql"
	"fmt"
)

const (
	dbport = 1605
)

var dbSingleton *sql.DB = nil

func GetDB() *sql.DB {
	if dbSingleton == nil {
		connStr := fmt.Sprintf("postgres://dbdata:dbpswd@localhost:%d/lesson?sslmode=disable", dbport)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		dbSingleton = db
	}

	return dbSingleton
}
