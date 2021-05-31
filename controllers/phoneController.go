package phoneController

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lcaa92/list_phone_jumia/models"
	_ "github.com/mattn/go-sqlite3"
)

type PhoneFilter struct {
	Country string
	State   string
}

// ListPhone return list of Phone
func ListPhone(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	phoneFilter := PhoneFilter{
		Country: strings.Join(query["country"], ""),
		State:   strings.Join(query["state"], ""),
	}

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
	var arrResult []models.PhoneResult

	for rows.Next() {
		err = rows.Scan(&phoneNumber)
		if err != nil {
			panic(err)
		}

		phoneModel := models.Phone{
			PhoneNumber: phoneNumber,
		}
		phoneModel.ExtractPhoneData()

		if phoneFilter.Country != "" && phoneFilter.Country != phoneModel.OutputPhone.Country {
			continue
		}

		if phoneFilter.State != "" && phoneFilter.State != phoneModel.OutputPhone.State {
			continue
		}
		arrResult = append(arrResult, phoneModel.OutputPhone)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(arrResult)
	if err != nil {
		panic(err)
	}
}
