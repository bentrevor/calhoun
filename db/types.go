package db

import "mime/multipart"

type QueryOpts struct {
	User  User
	Photo Photo
}

// e.g. postgres vs. in-memory
// the PhotoStore will use this to get a list of photo ids to use (queries will eventually get more
// complicated than `WHERE user_id = %d`)
type PhotoOrganizer interface {
	Insert(QueryOpts) int
	Select(QueryOpts) []Photo
}

// e.g. filesystem vs. S3
// I don't really need ReadPhoto() for now, since I just need PhotoId -> <img> tag.  The requests
// they make hit the FileServer, which handles reading
type PhotoPersister interface {
	WritePhoto(Photo)
	CountPhotos() int // mostly for debugging/testing
}

type User struct {
	Id   int64
	Name string
}

type Photo struct {
	Id        int
	PhotoFile *multipart.File
}
