package main

import (
	"flag"
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

	var dir string
	flag.StringVar(&dir, "static", "./static", "./static")
	flag.Parse()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./templates/")))

	log.Fatal(http.ListenAndServe(":8080", router))
}
