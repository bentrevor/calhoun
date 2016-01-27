package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	db.Exec(fmt.Sprintf("insert into users (id, name) values (%d, '%s')", user.Id, user.Name))
}

func OptsToPostgres(statementType StatementType, opts QueryOpts) string {
	switch statementType {
	case InsertStatement:
		columns := "filepath, user_id"
		values := fmt.Sprintf("'%s', %d", opts.Photo.Filepath, opts.User.Id)
		stmt := fmt.Sprintf("INSERT INTO photos (%s) VALUES (%s)", columns, values)
		fmt.Printf("\ninsert statement: %s\n\n", stmt)
		return stmt

	case SelectStatement:
		return fmt.Sprintf("SELECT filepath FROM photos WHERE user_id = %d;", opts.User.Id)
	}

	return ""
}

func NewPostgresTestDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=calhoun_admin dbname=calhoun_test sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	db.Exec("DELETE FROM photos *;")
	return &PostgresDB{DB: db}
}

func (db *PostgresDB) Insert(opts QueryOpts) (bool, error) {
	db.Exec(OptsToPostgres(InsertStatement, opts))

	return true, nil
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
		var filepath string
		err = rows.Scan(&filepath)

		photos = append(photos, Photo{Filepath: filepath})
	}

	if err != nil {
		log.Fatal("\n\n\ngot an error scanning db row:\n  ", err, "\n\n\n")
	}

	return photos
}
