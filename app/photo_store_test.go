package app_test

import (
	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPhotoStore_Filepath(t *testing.T) {
	// photos will be stored in /srv/images, in directories based on the md5 hash of their id
	// padded in front with 0's to 12 decimal places

	// $ echo -n "000000000012" | md5sum
	// 9ed63b492437de85736cb562f91f203c  -
	want := "/fake/srv/path/9e/d6/3b492437de85736cb562f91f203c"

	Describe("image filepaths")
	photo := Photo{Id: 12}
	memoryDB := NewMemoryDB()
	store := PhotoStore{DB: &memoryDB, SrvPath: "/fake/srv/path"}

	It("takes the md5 hash of (photo_id padded in front with 0s to 12 places)")
	AssertEquals(t, want, store.PhotoFilepath(photo))
}

func TestPhotoStore_SavingPhotos(t *testing.T) {
	Describe("PhotoStore")
	memoryDB := NewMemoryDB()
	memoryFS := NewMemoryFS()
	user := User{Name: "ben"}
	photo := Photo{Id: 1}
	store := PhotoStore{DB: &memoryDB, FS: &memoryFS, SrvPath: "/fake/srv/path"}

	store.SavePhoto(user, photo.PhotoFile)

	It("stores it in the db")
	AssertEquals(t, 1, len(store.PhotosForUser(user)))

	It("saves the file to the filesystem")
	AssertEquals(t, 1, store.FS.CountPhotos())
}
