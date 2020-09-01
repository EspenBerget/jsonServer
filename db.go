package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func connect(name string) *sql.DB {
	conn, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal("Could not open db.")
	}
	stmt, err := conn.Prepare(
		`CREATE TABLE IF NOT EXISTS ideas (
			id INTEGER PRIMARY KEY,
			ide TEXT NOT NULL,
			idea TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatal("MAKE STATEMENT: ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("EXECUTE STATEMENT: ", err)
	}

	return conn
}

func add(db *sql.DB, s Suggestion) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO ideas(ide,idea) VALUES (?, ?)")
	if err != nil {
		log.Println("INSERT ERROR:", err)
		return 0, err
	}
	r, err := stmt.Exec(s.IDE, s.Idea)
	if err != nil {
		log.Println("INSERT ERROR:", err)
		return 0, err
	}

	return r.LastInsertId()
}

func getAll(db *sql.DB) ([]Suggestion, error) {
	rows, err := db.Query("SELECT * FROM ideas")
	if err != nil {
		log.Println("QUERY ERROR:", err)
		return nil, err
	}
	ss := make([]Suggestion, 0, 16)
	for rows.Next() {
		s := Suggestion{}
		err = rows.Scan(&s.ID, &s.IDE, &s.Idea)
		if err != nil {
			log.Println("QUERY ERROR:", err)
			continue
		}
		ss = append(ss, s)
	}
	rows.Close()
	return ss, nil
}

func getID(db *sql.DB, id int64) (Suggestion, error) {
	row := db.QueryRow("SELECT * FROM ideas WHERE id=?", id)
	s := Suggestion{}
	err := row.Scan(&s.ID, &s.IDE, &s.Idea)
	if err != nil {
		log.Println("QUERY ERROR:", err)
		return s, err // Can't return nil here... WHY?
	}
	return s, nil
}

func delete(db *sql.DB, id int64) error {
	stmt, err := db.Prepare("DELETE FROM ideas WHERE id=?")
	if err != nil {
		log.Println("DELETE ERROR:", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("DELETE ERROR:", err)
		return err
	}
	return nil
}

// TODO update
