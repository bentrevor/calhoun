package app

import (
	"crypto/md5"
	"fmt"

	. "github.com/bentrevor/calhoun/db"
)

type PhotoStore struct {
	DB PhotoDB
}

func (store PhotoStore) SavePhoto(user User, photo Photo) (bool, error) {
	options := QueryOpts{User: user, Photo: photo}

	store.DB.Insert(options)

	return true, nil
}

func (store PhotoStore) PhotosForUser(user User) []Photo {
	options := QueryOpts{User: user}

	return store.DB.Select(options)
}

func (store PhotoStore) PhotoFilepath(photo Photo) string {
	paddedId := fmt.Sprintf("%012d", photo.Id)
	imgMD5 := md5.Sum([]byte(paddedId))
	hashedImgLocation := fmt.Sprintf("%x/%x/%x",
		imgMD5[0],
		imgMD5[1],
		imgMD5[2:],
	)

	return fmt.Sprintf("/srv/images/%s", hashedImgLocation)
}
