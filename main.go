package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	. "github.com/bentrevor/calhoun/app"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "some request info:\n%q\n%q\n%q\n%q",
		html.EscapeString(r.URL.Path),
		html.EscapeString(r.Host),
		html.EscapeString(r.RequestURI),
		html.EscapeString(r.RemoteAddr),
	)
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/upload_photo", UploadFormHandler)
	http.HandleFunc("/upload", UploadHandler)

	log.Print("server starting on 8080...\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
