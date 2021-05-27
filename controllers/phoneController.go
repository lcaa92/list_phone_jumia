package phoneController

import (
	"fmt"
	"net/http"
)

// ListPhone return list of Phone
func ListPhone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List Phone")
	fmt.Fprintln(w, "List Phones!")
}
