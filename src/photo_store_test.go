package calhoun_test

import (
	. "github.com/bentrevor/calhoun/src"
	. "github.com/bentrevor/calhoun/src/db"

	"testing"

	. "github.com/bentrevor/calhoun/src/spec-helper"
)

func TestPhotoStore_CanSavePhoto(t *testing.T) {
	Describe("PhotoStore")
	memoryDB := NewMemoryDB()
	user := User{Name: "ben"}
	photo := Photo{Filepath: "/path/to/file"}
	store := PhotoStore{DB: &memoryDB}

	store.SavePhoto(user, photo)

	It("stores it in the db")
	AssertEquals(t, 1, len(store.PhotosForUser(user)))
}
