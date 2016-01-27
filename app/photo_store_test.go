package app_test

import (
	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
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
