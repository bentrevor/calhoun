package calhoun

import (
	"database/sql"
	"log"
)

type PostgresDB struct {
	*sql.DB
}

func NewPostgresDB() *PostgresDB {
	// TODO don't disable ssl...
	// TODO environment variables for db name
	db, err := sql.Open("postgres", "user=calhoun_admin dbname=calhoun_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{DB: db}
}
