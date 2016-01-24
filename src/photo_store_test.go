package calhoun_test

import (
	. "github.com/bentrevor/calhoun/src"

	"testing"
)

func TestPhotoStore_CanSavePhoto(t *testing.T) {
	describe("saving photo")
	memoryDB := MemoryDB{}
	user := User{Name: "ben"}
	photo := Photo{Filepath: "/path/to/file"}
	store := PhotoStore{DB: &memoryDB}

	store.SavePhoto(user, photo)

	it("stores it in the db")
	assertEquals(t, 1, len(store.PhotosForUser(user)))
}
