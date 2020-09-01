package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestJsonGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jsonGetAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: want %v got %v", http.StatusOK, status)
	}
}

func TestJsonAdd(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{"ide":"vscode","idea":"more fun"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jsonAdd)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: want %v got %v", http.StatusOK, status)
	}

	// TODO add response test
}

func TestJsonGetID(t *testing.T) {
	req, err := http.NewRequest("GET", "/get/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/get/{id:[0-9]+}", jsonGetID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: want %v got %v", http.StatusOK, status)
	}

	expected := `{"id":1,"ide":"vscode","idea":"purescript support"}`
	if rr.Body.String() != expected {
		t.Errorf("Wrong body returned: want %v got %v", expected, rr.Body.String())
	}
}

func TestJsonDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/delete/{id:[0-9]+}", jsonDelete)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: want %v got %v", http.StatusOK, status)
	}

	expected := "Id 5 deleted"
	if rr.Body.String() != expected {
		t.Errorf("Wrong body returned: want %v got %v", expected, rr.Body.String())
	}
}

// TODO test if handlers fail correctly
