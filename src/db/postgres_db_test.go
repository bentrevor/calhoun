package db_test

import (
	// . "github.com/bentrevor/calhoun/src/db"

	"testing"

	. "github.com/bentrevor/calhoun/src/spec-helper"
)

// TODO shared examples?
func TestPhotoDB_PostgresDB(t *testing.T) {
	t.Skip("no!!!!!!!!!!!!!!!!")
	Describe("postgres db")

	// postgresDB := NewPostgresDB()

	It("starts empty")
	AssertEquals(t, 0, 2)

	// it("can insert a photo")
	// photo := Photo{Filepath: "picture"}
	// user := User{Name: "the user"}
	// otherUser := User{Name: "someone else"}

	// postgresDB.Insert(photo, QueryOpts{"user": user.Name})

	// assertEquals(t, 1, len(postgresDB.Photos))
	// assertEquals(t, 1, len(postgresDB.Select(QueryOpts{"user": user.Name})))
	// assertEquals(t, 0, len(postgresDB.Select(QueryOpts{"user": otherUser.Name})))

	// it("can select photos")
	// assertEquals(t, []Photo{photo}, postgresDB.Select(QueryOpts{"user": user.Name}))
}
