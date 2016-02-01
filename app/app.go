package app

import (
	"fmt"
	"io"
	"os"
)

type CalhounApp interface {
	UploadPhoto(io.Writer, *os.File)
	UploadPhotoForm(io.Writer)
	ViewPhotos(io.Writer, []Photo)
}

type Calhoun struct {
	Store    CalhounStore
	Renderer CalhounRenderer
}

// "domain models"
type User struct {
	Id   int64
	Name string
}

type Photo struct {
	Id        int
	PhotoFile *os.File
}

type CalhounFile struct {
}

func Run(environment string, server CalhounServer) {
	server.RegisterRoutes()

	if environment != "test" {
		server.Start()
	}
}

func (c Calhoun) UploadPhoto(w io.Writer, file *os.File) {
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
