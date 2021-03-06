package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	. "github.com/bentrevor/calhoun/app"
)

type PostgresDB struct {
	*sql.DB
}

type StatementType int

const (
	InsertStatement StatementType = iota
	SelectStatement
)

func (db *PostgresDB) InsertUser(user User) {
	db.Exec(fmt.Sprintf("insert into users (name) values ('%s')", user.Name))
}

func OptsToPostgres(statementType StatementType, opts QueryOpts) string {
	switch statementType {
	case InsertStatement:
		columns := "user_id"
		values := fmt.Sprintf("%d", opts.User.Id)
		return fmt.Sprintf("INSERT INTO photos (%s) VALUES (%s) RETURNING id;", columns, values)

	case SelectStatement:
		return fmt.Sprintf("SELECT id FROM photos WHERE user_id = %d;", opts.User.Id)
	}

	return ""
}

func NewPostgresTestDB() *PostgresDB {
	db := NewPostgresDB("test")
	db.Exec("DELETE FROM photos *;") // "database cleaner"
	db.Exec("DELETE FROM users *;")
	db.Exec("ALTER SEQUENCE photos_id_seq RESTART WITH 1;")
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1;")
	return db
}

func NewPostgresDB(environment string) *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", fmt.Sprintf("user=calhoun_admin dbname=calhoun_%s sslmode=disable", environment))

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{DB: db}
}

func (db *PostgresDB) Insert(opts QueryOpts) int {
	var photoId int
	query := OptsToPostgres(InsertStatement, opts)

	err := db.QueryRow(query).Scan(&photoId)
	if err != nil {
		log.Fatal("failure inserting into database: ", err)
	}

	return photoId
}

func (db *PostgresDB) Select(opts QueryOpts) []Photo {
	query := OptsToPostgres(SelectStatement, opts)
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		log.Fatal("\n\n\ngot an error SELECTing photos:\n  ", query, "\nerr:\n  ", err, "\n\n\n")
	}

	photos := []Photo{}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		photos = append(photos, Photo{Id: id})
	}

	if err != nil {
		log.Fatal("\n\n\ngot an error scanning db row:\n  ", err, "\n\n\n")
	}

	return photos
}
