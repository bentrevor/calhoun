package app_test

import (
	. "github.com/bentrevor/calhoun/app"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestCalhounStore_Creation(t *testing.T) {
	Describe("CalhounStore: creation")

	It("injects the FS, DB, and SrvPath")
	store := CalhounStore{FS: NewMemoryFS(), DB: NewMemoryDB(), SrvPath: "/fake/srv/path"}
	AssertEquals(t, "/fake/srv/path", store.SrvPath)
}

func TestCalhounStore_SavingPhotos(t *testing.T) {
	Describe("CalhounStore: saving photos")
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	store := CalhounStore{FS: NewMemoryFS(), DB: NewMemoryDB(), SrvPath: "/fake/srv/path"}

	store.SavePhoto(user, photo.PhotoFile)

	It("stores it in the db")
	AssertEquals(t, 1, len(store.PhotosForUser(user)))

	It("saves the file to the filesystem")
	AssertEquals(t, 1, store.FS.CountPhotos())
}
