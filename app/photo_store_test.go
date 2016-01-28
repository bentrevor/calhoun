package app_test

import (
	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPhotoStore_Creation(t *testing.T) {
	Describe("PhotoStore: creation")

	It("uses a fake srv path for the test env")
	store := NewPhotoStore("test", "")
	AssertEquals(t, "/fake/srv/path", store.SrvPath)

	It("uses a PhotoFS rooted at the injected srvPath for the dev env")
	storeB := NewPhotoStore("dev", "/injected/filepath")
	AssertEquals(t, "/injected/filepath", storeB.SrvPath)
	AssertEquals(t, "/injected/filepath", storeB.FS.RootDir())
}

func TestPhotoStore_SavingPhotos(t *testing.T) {
	Describe("PhotoStore: saving photos")
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	store := NewPhotoStore("test", "/fake/srv/path")

	store.SavePhoto(user, photo.PhotoFile)

	It("stores it in the db")
	AssertEquals(t, 1, len(store.PhotosForUser(user)))

	It("saves the file to the filesystem")
	AssertEquals(t, 1, store.FS.CountPhotos())
}
