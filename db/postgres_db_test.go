package db_test

import (
	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPostgresDB_OptsToPostgresInsert(t *testing.T) {
	Describe("PostgresDB: building INSERT query")
	It("uses columns/values")

	user := User{Id: 1, Name: "user asdf"}
	photo := Photo{Id: 2}

	want := "INSERT INTO photos (user_id) VALUES (1) RETURNING id;"
	got := OptsToPostgres(InsertStatement, QueryOpts{User: user, Photo: photo})

	AssertEquals(t, want, got)
}

func TestPostgresDB_OptsToPostgresSelect(t *testing.T) {
	Describe("PostgresDB: building SELECT query")
	It("can only `SELECT id FROM photos` by user_id") // MVP!

	user := User{Id: 1, Name: "user asdf"}

	want := "SELECT id FROM photos WHERE user_id = 1;"
	got := OptsToPostgres(SelectStatement, QueryOpts{User: user})

	AssertEquals(t, want, got)
}

// TODO shared examples?
func TestPostgresDB_SavingPhotos(t *testing.T) {
	Describe("PostgresDB: saving photos")
	postgresDB := NewPostgresTestDB()

	It("can insert a photo")
	user := User{Name: "the user", Id: 1}
	otherUser := User{Name: "someone else", Id: 2}

	postgresDB.InsertUser(user)
	postgresDB.InsertUser(otherUser)

	firstPhotoId := postgresDB.Insert(QueryOpts{User: user})

	AssertEquals(t, 1, len(postgresDB.Select(QueryOpts{User: user})))
	AssertEquals(t, 0, len(postgresDB.Select(QueryOpts{User: otherUser})))

	It("returns the id (for figuring out filesystem name)")
	secondPhotoId := postgresDB.Insert(QueryOpts{User: user})
	AssertEquals(t, firstPhotoId+1, secondPhotoId)

	It("can select photos")
	photos := []Photo{
		Photo{Id: firstPhotoId},
		Photo{Id: secondPhotoId},
	}
	AssertEquals(t, photos, postgresDB.Select(QueryOpts{User: user}))
}
