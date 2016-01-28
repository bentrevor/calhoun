package app

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"

	. "github.com/bentrevor/calhoun/db"
)

type PhotoStore struct {
	FS      PhotoPersister
	DB      PhotoOrganizer
	SrvPath string
}

func (store PhotoStore) SavePhoto(user User, photoFile *multipart.File) (bool, error) {
	photoId := store.savePhotoToDB(user)
	store.savePhotoToFS(photoFile, photoId)

	return true, nil
}

func (store PhotoStore) savePhotoToDB(user User) int {
	return store.DB.Insert(QueryOpts{User: user})
}

func (store PhotoStore) savePhotoToFS(photoFile *multipart.File, photoId int) {
	photo := Photo{Id: photoId, PhotoFile: photoFile}
	store.FS.WritePhoto(photo)
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
