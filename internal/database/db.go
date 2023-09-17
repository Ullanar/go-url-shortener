package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"url-shortener/internal/config"
)

func New(database config.Database) *sql.DB {
	connString := "postgres://" + database.User + ":" + database.Password + "@" + database.Host + "/" + database.DB + "?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Unable to connect Database: %s", err)
	}
	return db
}
