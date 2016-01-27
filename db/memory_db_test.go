package db_test

import (
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPhotoDB_MemoryDB(t *testing.T) {
	Describe("MemoryDB")

	memoryDB := NewMemoryDB()

	It("starts as an empty map")
	AssertEquals(t, 0, len(memoryDB.Select(QueryOpts{})))

	It("can insert a photo")
	photo := Photo{Filepath: "picture"}
	user := User{Name: "the user"}
	otherUser := User{Name: "someone else"}

	memoryDB.Insert(photo, QueryOpts{"user": user.Name})

	AssertEquals(t, 1, len(memoryDB.Photos))
	AssertEquals(t, 1, len(memoryDB.Select(QueryOpts{"user": user.Name})))
	AssertEquals(t, 0, len(memoryDB.Select(QueryOpts{"user": otherUser.Name})))

	It("can select photos")
	AssertEquals(t, []Photo{photo}, memoryDB.Select(QueryOpts{"user": user.Name}))
}
