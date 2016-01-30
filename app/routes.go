package app

import (
	"fmt"
	"net/http"

	. "github.com/bentrevor/calhoun/db"
	. "github.com/bentrevor/calhoun/ui"
)

func UploadFormHandler(rootDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewsPath := fmt.Sprintf("%s/views", rootDir)
		htmlFilename := "upload_photo_form.html"
		renderer := BrowserRenderer{RespWriter: w, ViewsPath: viewsPath}

		renderer.RenderHtmlFile(htmlFilename)
	}
}

func UploadHandler(srvPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO figure out global variable for environment
		file, _, err := r.FormFile("photoUpload")
		defer file.Close()

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		db := NewPostgresDB("dev")
		store := PhotoStore{DB: db, SrvPath: srvPath}
		store.SavePhoto(User{Name: "ben"}, &file)
	}
}
