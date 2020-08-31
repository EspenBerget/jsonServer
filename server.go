package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func checkInternalError(err error) {
	if err != nil {
		panic(http.ErrAbortHandler)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func jsonGetID(w http.ResponseWriter, r *http.Request) {
	vs := mux.Vars(r)
	idString := vs["id"]
	id, err := strconv.Atoi(idString)
	checkInternalError(err)

	s, err := getID(db, int64(id))
	checkInternalError(err)

	b, err := json.Marshal(s)
	checkInternalError(err)

	fmt.Fprint(w, string(b))
}

func jsonGetAll(w http.ResponseWriter, r *http.Request) {
	ss, err := getAll(db)
	checkInternalError(err)

	b, err := json.Marshal(ss)
	checkInternalError(err)

	fmt.Fprint(w, string(b))
}

func jsonAdd(w http.ResponseWriter, r *http.Request) {
	s := Suggestion{}
	b := make([]byte, 1024)
	n, err := r.Body.Read(b)
	if err != io.EOF {
		checkInternalError(err)
	}

	b = b[:n]
	checkInternalError(json.Unmarshal(b, &s))

	id, err := add(db, s)
	checkInternalError(err)

	fmt.Fprintf(w, "Added idea '%s' for IDE '%s' with ID %d", s.Idea, s.IDE, id)
}

func jsonDelete(w http.ResponseWriter, r *http.Request) {
	vs := mux.Vars(r)
	idString := vs["id"]
	id, err := strconv.Atoi(idString)
	checkInternalError(err)

	checkInternalError(delete(db, int64(id)))

	fmt.Fprintf(w, "Id %d deleted", id)
}
