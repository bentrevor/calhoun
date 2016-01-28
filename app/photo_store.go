package app

import (
	"crypto/md5"
	"fmt"

	. "github.com/bentrevor/calhoun/db"
)

type PhotoStore struct {
	DB      PhotoDB
	SrvPath string
}

func (store PhotoStore) SavePhoto(user User, photo Photo) (bool, error) {
	return store.DB.Insert(QueryOpts{User: user, Photo: photo})
}

func (store PhotoStore) PhotosForUser(user User) []Photo {
	return store.DB.Select(QueryOpts{User: user})
}

func (store PhotoStore) PhotoFilepath(photo Photo) string {
	paddedId := fmt.Sprintf("%012d", photo.Id)
	imgMD5 := md5.Sum([]byte(paddedId))
	hashedImgLocation := fmt.Sprintf("%x/%x/%x",
		imgMD5[0],
		imgMD5[1],
		imgMD5[2:],
	)

	return fmt.Sprintf("%s/%s", store.SrvPath, hashedImgLocation)
}
