package app

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "some request info:\n%q\n%q\n%q\n%q",
		html.EscapeString(r.URL.Path),
		html.EscapeString(r.Host),
		html.EscapeString(r.RequestURI),
		html.EscapeString(r.RemoteAddr),
	)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body := ""

	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		body = string(bytes)
	}

	fmt.Fprintf(w, "You said: %q", html.EscapeString(body))
}

// from https://www.socketloop.com/tutorials/golang-upload-file
// func UploadHandler(w http.ResponseWriter, r *http.Request) {

// 	// the FormFile function takes in the POST input id file
// 	file, header, err := r.FormFile("file")

// 	if err != nil {
// 		fmt.Fprintln(w, err)
// 		return
// 	}

// 	defer file.Close()

// 	out, err := os.Create("/tmp/uploadedfile")
// 	if err != nil {
// 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
// 		return
// 	}

// 	defer out.Close()

// 	// write the content from POST to the file
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		fmt.Fprintln(w, err)
// 	}

// 	fmt.Fprintf(w, "File uploaded successfully : ")
// 	fmt.Fprintf(w, header.Filename)
// }
