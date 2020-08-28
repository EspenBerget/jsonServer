package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func jsonGetID(w http.ResponseWriter, r *http.Request) {
	vs := mux.Vars(r)
	idString := vs["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		em := NewError("user error", "Misformed request id, must be a number.")
		fmt.Fprint(w, em)
		return
	}
	s, err := getID(db, int64(id))
	if err != nil {
		em := NewError("internal", err.Error())
		fmt.Fprint(w, em)
		return
	}
	b, err := json.Marshal(s)
	if err != nil {
		em := NewError("internal", err.Error())
		fmt.Fprint(w, em)
		return
	}
	fmt.Fprint(w, string(b))
}

func jsonGetAll(w http.ResponseWriter, r *http.Request) {
	ss, err := getAll(db)
	if err != nil {
		em := NewError("internal", err.Error())
		fmt.Fprint(w, em)
		return
	}
	b, err := json.Marshal(ss)
	if err != nil {
		em := NewError("internal", err.Error())
		fmt.Fprint(w, em)
		return
	}

	fmt.Fprint(w, string(b))
}

func jsonAdd(w http.ResponseWriter, r *http.Request) {
	s := Suggestion{}
	b := make([]byte, 1024)
	n, err := r.Body.Read(b)
	if err != io.EOF && err != nil {
		em := NewError("internal", "Error reading body")
		log.Println(err)
		fmt.Fprint(w, em)
		return
	}
	b = b[:n]
	if err := json.Unmarshal(b, &s); err != nil {
		em := NewError("internal", "Error parsing data in body")
		log.Println(err)
		fmt.Fprint(w, em)
		return
	}
	id, err := add(db, s)
	if err != nil {
		em := NewError("internal", "Error adding data")
		log.Println(err)
		fmt.Fprint(w, em)
		return
	}

	fmt.Fprintf(w, "Added idea '%s' for IDE '%s' with ID %d", s.Idea, s.IDE, id)
}

func jsonDelete(w http.ResponseWriter, r *http.Request) {
	vs := mux.Vars(r)
	idString := vs["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		em := NewError("user error", "Misformed request id, must be a number.")
		fmt.Fprint(w, em)
		return
	}
	if err := delete(db, int64(id)); err != nil {
		em := NewError("internal", fmt.Sprintf("Failed to delete id: %d, error message: %v", id, err))
		fmt.Fprint(w, em)
		return
	}

	fmt.Fprintf(w, "Id %d deleted", id)
}
