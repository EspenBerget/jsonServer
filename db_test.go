package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewSuggestion("vscode", "purescript support")
	_, err := add(db, s)
	if err != nil {
		t.Errorf("Add failed with %v", err)
	}
}

func TestDelete(t *testing.T) {
	s := NewSuggestion("blabla", "foo bar baz")
	id, err := add(db, s)
	if err != nil {
		t.Errorf("Add failure while testing delete: %v", err)
	}
	err = delete(db, id)
	if err != nil {
		t.Errorf("Delete error %v", err)
	}
}

func TestGetID(t *testing.T) {
	s := NewSuggestion("foo", "hello world")
	id, err := add(db, s)
	if err != nil {
		t.Errorf("Add failure while testing getID: %v", err)
	}
	a, err := getID(db, id)
	if err != nil {
		t.Errorf("getID failure %v", err)
	}
	if a.IDE != s.IDE || a.Idea != s.Idea {
		t.Error("getID failure: Returned value does not match input value")
	}
}

func TestGetAll(t *testing.T) {
	_, err := getAll(db)
	if err != nil {
		t.Errorf("getAll failure: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	s := NewSuggestion("update", "this is an update")
	s.ID = 2
	err := update(db, s)
	if err != nil {
		t.Errorf("update failure %v", err)
	}
	a, err := getID(db, 2)
	if err != nil {
		t.Errorf("update failure while getting updated value. Error: %v", err)
	}
	if a.ID != s.ID || a.IDE != s.IDE || a.Idea != s.Idea {
		t.Error("Updated value not the same as input value")
	}
}
