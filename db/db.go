package db

import (
	"database/sql"
	"log"
)

func NewPostgreSqlStorage(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
