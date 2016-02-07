package app

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

type CalhounAction int

const (
	UploadPhoto CalhounAction = iota
	UploadPhotoForm
	ViewPhotos
)

type Calhouner interface {
	LookupHandler(CalhounAction) CalhounHandler
}

type Calhoun struct {
	Store    CalhounStore
	Renderer CalhounRenderer
}

// hmmm...
func (c Calhoun) LookupHandler(action CalhounAction) CalhounHandler {
	switch action {

	case UploadPhoto:
		return func(w io.Writer, r *CalhounRequest) {
			c.UploadPhoto(w, r.UploadFile)
		}
	case UploadPhotoForm:
		return func(w io.Writer, _ *CalhounRequest) {
			c.UploadPhotoForm(w)
		}
	case ViewPhotos:
		return func(w io.Writer, _ *CalhounRequest) {
			c.ViewPhotos(w)
		}
	default:
		log.Fatal("unknown action: ", action)
	}

	return nil
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
		c.Renderer.Render(UploadPhoto, w)
	} else {
		fmt.Fprintln(w, "error saving photo: ", err)
	}
}

func (c Calhoun) UploadPhotoForm(w io.Writer) {
	c.Renderer.Render(UploadPhotoForm, w)
}

func (c Calhoun) ViewPhotos(w io.Writer) {
	user := User{Id: 1, Name: "God"} // until auth middleware is implemented
	photos := c.Store.PhotosForUser(user)

	c.Renderer.Render(ViewPhotos, w, RenderArgs{Photos: photos})
}
