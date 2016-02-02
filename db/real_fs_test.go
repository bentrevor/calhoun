package db_test

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestRealFS_Filepath(t *testing.T) {
	Describe("RealFS: filepath")
	// photos will be stored in <asset path>/images/srv, in directories based on the md5 hash of
	// their id padded in front with 0's to 12 decimal places

	// $ echo -n "000000000012" | md5sum
	// 9ed63b492437de85736cb562f91f203c  -
	want := "/fake/srv/path/9e/d6/3b492437de85736cb562f91f203c"

	photo := Photo{Id: 12}
	fs := RealFS{RootDir: "/fake/srv/path"}

	It("takes the md5 hash of (photo_id padded in front with 0s to 12 places)")
	AssertEquals(t, want, fs.PhotoFilepath(photo))
}

func TestRealFS_Writing(t *testing.T) {
	Describe("RealFS: writing file")
	rootDir := TestdataPath()
	photoPath := fmt.Sprintf("%s/images/dog.png", rootDir)

	photoFile, err := os.Open(photoPath)
	defer photoFile.Close()

	if err != nil {
		log.Fatal("in test: ", err)
	}

	// &multipart.File(photoFile) => cannot take the address of multipart.File(photoFile)
	mpPhoto := multipart.File(photoFile)
	photo := Photo{Id: 12, PhotoFile: &mpPhoto}

	fsRoot := fmt.Sprintf("%s/srv", rootDir)
	fs := RealFS{RootDir: fsRoot}
	fs.WritePhoto(photo)

	It("saves to the filesystem")

	_, err = os.Stat(fs.PhotoFilepath(photo))
	Assert(t, !os.IsNotExist(err))

	// TODO defer these
	os.RemoveAll(fsRoot)
	os.Mkdir(fsRoot, 0755)
}

func TestRealFS_Errors(t *testing.T) {
	Describe("RealFS: errors")
	rootDir := TestdataPath()
	photoPath := fmt.Sprintf("%s/images/dog.png", rootDir)

	photoFile, err := os.Open(photoPath)
	defer photoFile.Close()

	if err != nil {
		log.Fatal("in test 2: ", err)
	}

	mpPhoto := multipart.File(photoFile)
	photo := Photo{Id: 12, PhotoFile: &mpPhoto}

	fsRoot := fmt.Sprintf("%s/srv", rootDir)
	os.RemoveAll(fsRoot)
	os.Mkdir(fsRoot, 0755)

	fs := RealFS{RootDir: fsRoot}

	It("raises an error when the filepath already exists")
	// this is actually pretty likely to happen - my filesystem and database sequence have to be
	// synchronized, so it's only a matter of time before they get out of sync

	_ = fs.WritePhoto(photo)
	err = fs.WritePhoto(photo)

	AssertNotEquals(t, err, nil)

	os.RemoveAll(fsRoot)
	os.Mkdir(fsRoot, 0755)
}
