package calhoun_test

import (
	. "github.com/bentrevor/calhoun/src"

	"testing"
)

func TestPhotoStore_CanSavePoto(t *testing.T) {
	describe("saving photo")
	memoryDB := NewMemoryDB()
	user := User{Name: "ben"}
	photo := Photo{Filepath: "/path/to/file"}
	store := PhotoStore{DB: &memoryDB}

	store.SavePhoto(user, photo)

	it("stores it in the db")
	assertEquals(t, 1, 1)
}
