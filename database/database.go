package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetPhoneList() (phoneList []string) {
	db, err := sql.Open("sqlite3", "./sample.db?mode=memory&_fk=true&cache=shared")
	defer db.Close()
	if err != nil {
		panic(err)
	}

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
