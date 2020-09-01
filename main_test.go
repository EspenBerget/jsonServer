package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database := "./test_ideas.db"
	db = connect(database)
	code := m.Run()

	err := os.Remove(database)
	if err != nil {
		log.Printf("Failed to delete database '%s' after tests\n", database)
	}
	os.Exit(code)
}
