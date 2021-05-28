package models

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
	p.OutputPhone.Country = "Brasil"
	p.OutputPhone.State = "OK"
	p.OutputPhone.CountryCode = "55"
	p.OutputPhone.PhoneNumber = p.PhoneNumber
}
