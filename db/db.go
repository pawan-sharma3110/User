package db

import (
	"database/sql"
	"fmt"

	"log"
	"user/model"
	"user/utils"

	"errors"

	"github.com/google/uuid"
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

func CreateUserTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}

	fmt.Println("User table created successfully")
}
func InsertUser(db *sql.DB, user model.User) (uuid.UUID, error) {
	// Hash the user's password
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return uuid.Nil, err
	}

	// Insert the user into the users table
	insertQuery := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	var id uuid.UUID
	err = db.QueryRow(insertQuery, user.Email, hashPass).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
