package main

// Suggestion struct represents one idea
type Suggestion struct {
	ID   int64  `json:"id"`
	IDE  string `json:"ide"`
	Idea string `json:"idea"`
}

// NewSuggestion Returns a new suggestion with ID set to 0
func NewSuggestion(ide, idea string) Suggestion {
	return Suggestion{ID: 0, IDE: ide, Idea: idea}
}
