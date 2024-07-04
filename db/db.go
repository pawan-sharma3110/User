package db

import (
	"database/sql"

	"log"
	"user/model"

	"errors"

	_ "github.com/lib/pq"
)

func DbIn() (db *sql.DB, err error) {
	conStr := `host=localhost port=5432 user=postgres password=Pawan@2003 dbname=user sslmode=disable`
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		err = errors.New("database connection error")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		err = errors.New("database connection still alive")

		return nil, err
	}
	return db, nil
}

func InsertUser(db *sql.DB, user model.User) error {

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}

	query = `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, user.Id, user.Email, user.Password)
	return err
}
