package db

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	. "github.com/bentrevor/calhoun/app"
)

type RealFS struct {
	RootDir string
}

func (fs RealFS) WritePhoto(photo Photo) error {
	photoFilepath := fs.PhotoFilepath(photo)
	dirs := strings.Split(photoFilepath, "/")
	photoDir := strings.Join(dirs[:len(dirs)-1], "/")

	err := os.MkdirAll(photoDir, 0755)

	if err != nil {
		log.Print("Unable to create the photo dir:  ", err)
		return err
	}

	if fs.photoAlreadyExists(photoFilepath) {
		return errors.New("photo with that id already saved, something went wrong...")
	}

	out, err := os.Create(photoFilepath)
	defer out.Close()

	if err != nil {
		log.Fatal("Unable to create the file for writing:  ", err)
		return err
	}

	_, err = io.Copy(out, *photo.PhotoFile)
	if err != nil {
		log.Fatal("Unable to copy photo file for photo #", photo.Id, ":  ", err)
		return err
	}
	return nil
}

func (fs RealFS) photoAlreadyExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)

}

func (fs RealFS) CountPhotos() int {
	return 50 // TODO count files in srv directory
}

// This might eventually go in the app package, since the presentation layer needs it too, but I'll
// wait until I actually implement a new FS for that
func (fs RealFS) PhotoSrc(id int) string {
	paddedId := fmt.Sprintf("%012d", id)
	imgMD5 := md5.Sum([]byte(paddedId))

	return fmt.Sprintf("%x/%x/%x",
		imgMD5[0],
		imgMD5[1],
		imgMD5[2:],
	)
}

func (fs RealFS) PhotoFilepath(photo Photo) string {
	return fmt.Sprintf("%s/%s", fs.RootDir, fs.PhotoSrc(photo.Id))
}
