package app

import (
	"fmt"
	"io"
	"mime/multipart"
)

// TODO get rid of ISP violations
// type CalhounAction int
// const (
// 	UploadPhoto CalhounAction = iota
// 	UploadPhotoForm
// 	ViewPhotos
// )

type CalhounApp interface {
	UploadPhoto(io.Writer, *multipart.File)
	UploadPhotoForm(io.Writer)
	ViewPhotos(io.Writer)
}

type Calhoun struct {
	Store    CalhounStore
	Renderer CalhounRenderer
}

type User struct {
	Id   int64
	Name string
}

type Photo struct {
	Id        int
	PhotoFile *multipart.File
	Src       string
}

func Run(environment string, server CalhounServer) {
	server.RegisterRoutes()

	if environment != "test" {
		server.Start()
	}
}

func (c Calhoun) UploadPhoto(w io.Writer, file *multipart.File) {
	user := User{Id: 1, Name: "God"} // until auth middleware is implemented
	err := c.Store.SavePhoto(user, file)

	if err == nil {
		c.Renderer.UploadPhoto(w)
	} else {
		fmt.Fprintln(w, "error saving photo: ", err)
	}
}

func (c Calhoun) UploadPhotoForm(w io.Writer) {
	c.Renderer.UploadPhotoForm(w)
}

func (c Calhoun) ViewPhotos(w io.Writer) {
	user := User{Id: 1, Name: "God"} // until auth middleware is implemented
	photos := c.Store.PhotosForUser(user)
	c.Renderer.ViewPhotos(w, photos)
}
