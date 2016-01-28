package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/bentrevor/calhoun/db"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	htmlFilename := "views/upload_photo.html"
	body, err := ioutil.ReadFile(htmlFilename)

	if err != nil {
		fmt.Sprint("error")
	}

	fmt.Fprint(w, string(body))
}

func UploadHandler(srvPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO figure out global variable for environment
		db := NewPostgresDB("dev")
		store := PhotoStore{DB: db, SrvPath: srvPath}
		file, _, err := r.FormFile("photoUpload")

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer file.Close()

		store.SavePhoto(User{Name: "ben"}, &file)
	}
}
