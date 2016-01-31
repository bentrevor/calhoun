package app

import (
	"mime/multipart"
	"net/http"
)

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
	RootDir() string
	WritePhoto(Photo)
	CountPhotos() int // mostly for debugging/testing
}

type CalhounRenderer interface {
	Handle(string) http.HandlerFunc
}

type QueryOpts struct {
	User  User
	Photo Photo
}

type User struct {
	Id   int64
	Name string
}

type Photo struct {
	Id        int
	PhotoFile *multipart.File
}
