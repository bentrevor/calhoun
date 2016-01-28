package db_test

import (
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPhotoDB_OptsToPostgresInsert(t *testing.T) {
	Describe("building pg INSERT")
	It("uses columns/values")

	user := User{Id: 1, Name: "user asdf"}
	photo := Photo{Id: 2}

	want := "INSERT INTO photos (user_id) VALUES (1)"
	got := OptsToPostgres(InsertStatement, QueryOpts{User: user, Photo: photo})

	AssertEquals(t, want, got)
}

func TestPhotoDB_OptsToPostgresSelect(t *testing.T) {
	Describe("building pg SELECT")
	It("can only `SELECT id FROM photos` by user_id") // MVP!

	user := User{Id: 1, Name: "user asdf"}

	want := "SELECT id FROM photos WHERE user_id = 1;"
	got := OptsToPostgres(SelectStatement, QueryOpts{User: user})

	AssertEquals(t, want, got)
}

// TODO shared examples?
func TestPhotoDB_PostgresDB(t *testing.T) {
	Describe("PostgresDB")
	postgresDB := NewPostgresTestDB()

	It("can insert a photo")
	user := User{Name: "the user", Id: 1}
	otherUser := User{Name: "someone else", Id: 2}

	postgresDB.InsertUser(user)
	postgresDB.InsertUser(otherUser)

	photoId := postgresDB.Insert(QueryOpts{User: user})

	AssertEquals(t, 1, len(postgresDB.Select(QueryOpts{User: user})))
	AssertEquals(t, 0, len(postgresDB.Select(QueryOpts{User: otherUser})))

	It("can select photos")
	photo := Photo{Id: photoId}
	AssertEquals(t, []Photo{photo}, postgresDB.Select(QueryOpts{User: user}))
}
