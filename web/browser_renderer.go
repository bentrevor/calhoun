package web

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Yield string
}

type BrowserRenderer struct {
	ViewsPath string
}

func (br *BrowserRenderer) renderHtmlFile(filename string, writer http.ResponseWriter, page Page) {
	layoutPath := fmt.Sprintf("%s/layout.html", br.ViewsPath)
	filepath := fmt.Sprintf("%s/%s.html", br.ViewsPath, filename)

	tmpl, err := template.ParseFiles(layoutPath, filepath)

	if err != nil {
		log.Fatal("error on template.ParseFiles: ", err)
	}

	tmpl.Execute(writer, page)
}

func (br *BrowserRenderer) UploadPhoto(w io.Writer, file *os.File) {
	user := User{Id: 1, Name: "God"} // until auth middleware is implemented
	err := s.Store.SavePhoto(user, file)

	if err == nil {
		br.RenderHtmlFile("upload_success", w, Page{})
	} else {
		fmt.Fprintln(w, "error saving photo: ", err)
		return
	}
}

func (br *BrowserRenderer) UploadPhotoForm(w io.Writer) {
	fmt.Fprint(w, "this is the upload form")
}

func (br *BrowserRenderer) ViewPhotos(w io.Writer, photos []Photo) {
	fmt.Fprint(w, "this is the view photos")
}
