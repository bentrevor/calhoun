package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	htmlFilename := "views/upload_photo.html"
	body, err := ioutil.ReadFile(htmlFilename)

	if err != nil {
		fmt.Sprint("error")
	}

	fmt.Fprint(w, string(body))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("photoUpload")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	uploadedFilename := fmt.Sprintf("/srv/images/%s", "adf")
	out, err := os.Create(uploadedFilename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}
