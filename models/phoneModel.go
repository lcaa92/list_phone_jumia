package models

import (
	"regexp"
	"strings"

	"github.com/lcaa92/list_phone_jumia/database"
)

type Phone struct {
	PhoneNumber string
	OutputPhone PhoneResult
}

type PhoneResult struct {
	Country     string
	State       string
	CountryCode string
	PhoneNumber string
}

func (p *Phone) ExtractPhoneData() {
	country, state, country_code, phone_number := getDataFromPhone(p.PhoneNumber)
	p.OutputPhone.Country = country
	p.OutputPhone.State = state
	p.OutputPhone.CountryCode = country_code
	p.OutputPhone.PhoneNumber = phone_number
}

func getDataFromPhone(phone_number string) (country, state, country_code, phone string) {
	split_phone := strings.Split(phone_number, " ")

	country_code = split_phone[0]
	country_code = strings.Replace(country_code, "(", "", -1)
	country_code = strings.Replace(country_code, ")", "", -1)
	country_code = strings.Join([]string{"+", country_code}, "")

	phone = split_phone[1]

	var regexState string

	switch country_code {
	case "+237":
		country = "Cameroon"
		regexState = `\(237\)\ ?[2368]\d{7,8}$`
	case "+251":
		country = "Ethiopia"
		regexState = `\(251\)\ ?[1-59]\d{8}$`
	case "+212":
		country = "Morocco"
		regexState = `\(212\)\ ?[5-9]\d{8}$`
	case "+258":
		country = "Mozambique"
		regexState = `\(258\)\ ?[28]\d{7,8}$`
	case "+256":
		country = "Uganda"
		regexState = `\(256\)\ ?\d{9}$`
	}

	var validateState = regexp.MustCompile(regexState)
	if validateState.MatchString(phone_number) {
		state = "OK"
	} else {
		state = "NOK"
	}
	return
}

func GetPhoneList() (phoneList []string) {
	db := database.OpenConnection()
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
