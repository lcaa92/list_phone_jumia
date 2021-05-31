package phoneController

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lcaa92/list_phone_jumia/models"
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

	phoneList := models.GetPhoneList()

	var arrResult []models.PhoneResult
	for _, phoneNumber := range phoneList {
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
	err := json.NewEncoder(w).Encode(arrResult)
	if err != nil {
		panic(err)
	}
}
