package db

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

// func (pg *PostgresDB) Insert(photo Photo, opts QueryOpts) (bool, error) {
// 	pg.Query("INSERT INTO photos (%s) values (%s)", columns, values)

// 	return true, nil
// }

// func (pg *PostgresDB) Select(opts QueryOpts) []Photo {
// 	scope = opts["scope"] || "*"
// 	userId = opts["userId"]
// 	tags = opts["tags"]
// 	query = buildFrom(userId, tags, etc...)

// 	return pg.Query("SELECT %s FROM photos WHERE %s", scope, query)
// }
