package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	tls := flag.Bool("s", false, "use TLS")
	flag.Parse()
	db = connect("./ideas.db")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/all", jsonGetAll).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", jsonGetID).Methods("GET")
	r.HandleFunc("/add", jsonAdd).Methods("POST")
	r.HandleFunc("/delete/{id:[0-9]+}", jsonDelete).Methods("DELETE")
	r.HandleFunc("/update/{id:[0-9]+}", jsonUpdate).Methods("PUT")

	fmt.Println("Server startet on localhost:9854")
	if *tls {
		log.Fatal(http.ListenAndServeTLS(":9854", "cert.pem", "key.pem", r))
	} else {
		log.Fatal(http.ListenAndServe(":9854", r))
	}

	db.Close()
}
