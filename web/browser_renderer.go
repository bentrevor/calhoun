package web

import (
	"fmt"
	"html/template"
	"io"
	"log"

	"github.com/bentrevor/calhoun/app"
)

type Page struct {
	PhotoSrcs []string
}

type BrowserRenderer struct {
	ViewsPath  string
	PhotosPath string
}

func (br *BrowserRenderer) renderHtmlFile(filename string, writer io.Writer, page Page) {
	layoutPath := fmt.Sprintf("%s/layout.html", br.ViewsPath)
	filepath := fmt.Sprintf("%s/%s.html", br.ViewsPath, filename)

	tmpl, err := template.ParseFiles(layoutPath, filepath)

	if err != nil {
		log.Fatal("error on template.ParseFiles: ", err)
	}

	tmpl.Execute(writer, page)
}

func (br BrowserRenderer) UploadPhoto(w io.Writer) {
	br.renderHtmlFile("upload_success", w, Page{})
}

func (br BrowserRenderer) UploadPhotoForm(w io.Writer) {
	br.renderHtmlFile("upload_photo_form", w, Page{})
}

func (br BrowserRenderer) ViewPhotos(w io.Writer, photos []app.Photo) {
	srcs := []string{}
	for _, photo := range photos {
		srcs = append(srcs, fmt.Sprintf("%s/%s", br.PhotosPath, photo.Src))
	}
	br.renderHtmlFile("view_photos", w, Page{PhotoSrcs: srcs})
}
