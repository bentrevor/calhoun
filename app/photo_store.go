package app

import (
	"log"
	"mime/multipart"

	. "github.com/bentrevor/calhoun/db"
)

type PhotoStore struct {
	FS      PhotoPersister
	DB      PhotoOrganizer
	SrvPath string
}

func NewPhotoStore(environment, srvPath string) PhotoStore {
	switch environment {
	case "test":
		db := NewMemoryDB()
		fs := NewMemoryFS()
		return PhotoStore{
			DB:      db,
			FS:      fs,
			SrvPath: "/fake/srv/path",
		}
	case "dev":
		db := NewPostgresDB(environment)
		fs := NewPhotoFS(srvPath)
		return PhotoStore{
			DB:      db,
			FS:      fs,
			SrvPath: srvPath,
		}
	default:
		log.Fatal("unknown env: ", environment)
	}

	return PhotoStore{}
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
