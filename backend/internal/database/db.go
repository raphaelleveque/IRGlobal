package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname, sslmode string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	log.Printf("Attempting to connect to PostgreSQL database at %s:%s", host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	log.Printf("Successfully connected to PostgreSQL database")
	return db, nil
}
