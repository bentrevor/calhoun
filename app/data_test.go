package app_test

import (
	. "github.com/bentrevor/calhoun/app"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestCalhounStore_SavingPhotos(t *testing.T) {
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	srvPath := "/fake/srv/path"
	store := CalhounStore{FS: NewMemoryFS(srvPath), DB: NewMemoryDB(), SrvPath: srvPath}

	store.SavePhoto(user, photo.PhotoFile)
	photos := store.PhotosForUser(user)

	Describe("CalhounStore: saving photos")
	It("stores it in the db")
	AssertEquals(t, 1, len(photos))

	It("saves the file to the filesystem")
	AssertEquals(t, 1, store.FS.CountPhotos())

	Describe("CalhounStore: loading photos")
	It("adds the Photo.Src")
	AssertEquals(t, "/fake/srv/path/1", photos[0].Src)
}
