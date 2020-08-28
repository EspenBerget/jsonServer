package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	db = connect("./ideas.db")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/all", jsonGetAll).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", jsonGetID).Methods("GET")
	r.HandleFunc("/add", jsonAdd).Methods("POST")
	r.HandleFunc("/delete/{id:[0-9]+}", jsonDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9854", r))
}
