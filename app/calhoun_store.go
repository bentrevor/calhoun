package app

import "mime/multipart"

type CalhounStore struct {
	FS      CalhounFS
	DB      CalhounDB
	SrvPath string
}

func (store CalhounStore) SavePhoto(user User, photoFile *multipart.File) error {
	photoId := store.savePhotoToDB(user)
	store.savePhotoToFS(photoFile, photoId)

	// TODO return real errors
	return nil
}

func (store CalhounStore) savePhotoToDB(user User) int {
	return store.DB.Insert(QueryOpts{User: user})
}

func (store CalhounStore) savePhotoToFS(photoFile *multipart.File, photoId int) {
	photo := Photo{Id: photoId, PhotoFile: photoFile}
	store.FS.WritePhoto(photo)
}

func (store CalhounStore) PhotosForUser(user User) []Photo {
	return store.DB.Select(QueryOpts{User: user})
}
