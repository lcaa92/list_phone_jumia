package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./sample.db?mode=memory&_fk=true&cache=shared")
	if err != nil {
		panic(err)
	}
	return db
}
