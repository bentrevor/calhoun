package app_test

import (
	. "github.com/bentrevor/calhoun/app"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestCalhounStore_SavingPhotos(t *testing.T) {
	Describe("CalhounStore: saving photos")
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	srvPath := "/fake/srv/path"
	store := CalhounStore{FS: NewMemoryFS(srvPath), DB: NewMemoryDB(), SrvPath: srvPath}

	store.SavePhoto(user, photo.PhotoFile)

	It("stores it in the db")
	AssertEquals(t, 1, len(store.PhotosForUser(user)))

	It("saves the file to the filesystem")
	AssertEquals(t, 1, store.FS.CountPhotos())
}

func TestCalhounStore_LoadingPhotos(t *testing.T) {
	Describe("CalhounStore: loading photos")
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	srvPath := "/fake/srv/path"
	store := CalhounStore{FS: NewMemoryFS(srvPath), DB: NewMemoryDB(), SrvPath: srvPath}

	store.SavePhoto(user, photo.PhotoFile)
	photos := store.PhotosForUser(user)

	It("adds the Photo.Src")
	AssertEquals(t, "/fake/srv/path/1", photos[0].Src)
}
