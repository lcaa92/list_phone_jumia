package models

import (
	"fmt"
	"strings"
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
	state = "OK"
	split_phone := strings.Split(phone_number, " ")

	country_code = split_phone[0]
	country_code = strings.Replace(country_code, "(", "", -1)
	country_code = strings.Replace(country_code, ")", "", -1)
	country_code = strings.Join([]string{"+", country_code}, "")

	phone = split_phone[1]

	switch country_code {
	case "+237":
		country = "Cameroon"
	case "+251":
		country = "Ethiopia"
	case "+212":
		country = "Morocco"
	case "+258":
		country = "Mozambique"
	case "+256":
		country = "Uganda"

	}
	return
}
