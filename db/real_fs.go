package db

import (
	"crypto/md5"
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

func (fs RealFS) WritePhoto(photo Photo) {
	photoFilepath := fs.PhotoFilepath(photo)
	dirs := strings.Split(photoFilepath, "/")
	photoDir := strings.Join(dirs[:len(dirs)-1], "/")

	err := os.MkdirAll(photoDir, 0755)

	if err != nil {
		log.Fatal("Unable to create the photo dir:  ", err)
	}

	out, err := os.Create(photoFilepath)
	defer out.Close()

	if err != nil {
		log.Fatal("Unable to create the file for writing:  ", err)
	}

	_, err = io.Copy(out, *photo.PhotoFile)
	if err != nil {
		log.Fatal("Unable to copy photo file for photo #", photo.Id, ":  ", err)
	}
}

func (fs RealFS) CountPhotos() int {
	return 50 // TODO count files in srv directory
}

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
