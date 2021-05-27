package phoneController

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// ListPhone return list of Phone
func ListPhone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List Phone")

	db, err := sql.Open("sqlite3", "./sample.db?mode=memory&_fk=true&cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM customer;")
	if err != nil {
		panic(err)
	}

	var uid int
	var name string
	var phone string

	for rows.Next() {
		err = rows.Scan(&uid, &name, &phone)
		if err != nil {
			panic(err)
		}
		fmt.Println(uid)
		fmt.Println(name)
		fmt.Println(phone)
	}

	rows.Close()
	fmt.Fprintln(w, "List Phones!")
}
