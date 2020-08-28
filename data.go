package main

import (
	"encoding/json"
	"log"
)

// Suggestion struct represents one idea
type Suggestion struct {
	ID   int64  `json:"id"`
	IDE  string `json:"ide"`
	Idea string `json:"idea"`
}

// ErrorMsg is sent back to users if there is an error with the prossesing of their request
type ErrorMsg struct {
	Kind string `json:"kind"`
	Msg  string `json:"message"`
}

// NewSuggestion Returns a new suggestion with ID set to 0
func NewSuggestion(ide, idea string) Suggestion {
	return Suggestion{ID: 0, IDE: ide, Idea: idea}
}

// NewError returns a new error message as a json string
func NewError(kind, msg string) string {
	e := ErrorMsg{Kind: kind, Msg: msg}
	b, err := json.Marshal(e)
	if err != nil {
		log.Println("Marshalling error, while trying to create an error...")
		return `{"kind":"internal", "message":"Failed to make error message"}`
	}
	return string(b)
}
