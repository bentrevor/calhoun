package app

import "mime/multipart"

type CalhounStore struct {
	FS      CalhounFS
	DB      CalhounDB
	SrvPath string
}

// e.g. postgres vs. in-memory
// the CalhounStore will use this to get a list of photo ids to use (queries will eventually get more
// complicated than `WHERE user_id = %d`)
type CalhounDB interface {
	Insert(QueryOpts) int
	Select(QueryOpts) []Photo
}

// e.g. filesystem vs. S3
// I don't really need ReadPhoto() for now, since I just need PhotoId -> <img> tag.  The requests
// they make hit the FileServer, which handles reading
type CalhounFS interface {
	WritePhoto(Photo)
	PhotoSrc(int) string
	CountPhotos() int // mostly for debugging/testing
}

type QueryOpts struct {
	User  User
	Photo Photo
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
	photos := store.DB.Select(QueryOpts{User: user})

	for i := range photos {
		photos[i].Src = store.FS.PhotoSrc(photos[i].Id)
	}

	return photos
}
