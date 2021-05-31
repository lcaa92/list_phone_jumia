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

func GetPhoneList() (phoneList []string) {
	db := OpenConnection()
	defer db.Close()

	rows, err := db.Query("SELECT phone FROM customer;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var phoneNumber string
	for rows.Next() {
		err = rows.Scan(&phoneNumber)
		if err != nil {
			panic(err)
		}
		phoneList = append(phoneList, phoneNumber)
	}

	return
}
