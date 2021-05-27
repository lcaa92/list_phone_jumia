package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	phoneController "github.com/lcaa92/list_phone_jumia/controllers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting ...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/phones", phoneController.ListPhone)

	log.Fatal(http.ListenAndServe(":8080", router))
}
