package db

import (
	"io"
	"log"
	"os"
)

type PhotoFS struct {
	Photos []Photo
}

func (fs *PhotoFS) WritePhoto(photo Photo) {
	photoFilename := "TODO" // hash from photoId
	out, err := os.Create(photoFilename)

	if err != nil {
		log.Fatal("Unable to create the file for writing:  ", err)
	}

	defer out.Close()

	_, err = io.Copy(out, *photo.PhotoFile)
	if err != nil {
		log.Fatal("Unable to copy photo file for photo #", photo.Id, ":  ", err)
	}
}

func (fs *PhotoFS) CountPhotos() int {
	return len(fs.Photos)
}
