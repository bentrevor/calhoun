package calhoun_test

import (
	. "github.com/bentrevor/calhoun/src"

	"testing"
)

func TestPhotoDB_MemoryDB(t *testing.T) {
	describe("memory db")

	memoryDB := NewMemoryDB()

	it("starts as an empty map")
	assertEquals(t, 0, len(memoryDB.Select(QueryOpts{})))

	it("can insert a photo")
	photo := Photo{Filepath: "picture"}
	user := User{Name: "the user"}
	otherUser := User{Name: "someone else"}

	memoryDB.Insert(photo, QueryOpts{"user": user.Name})

	assertEquals(t, 1, len(memoryDB.Photos))
	assertEquals(t, 1, len(memoryDB.Select(QueryOpts{"user": user.Name})))
	assertEquals(t, 0, len(memoryDB.Select(QueryOpts{"user": otherUser.Name})))

	it("can select photos")
	assertEquals(t, []Photo{photo}, memoryDB.Select(QueryOpts{"user": user.Name}))
}
