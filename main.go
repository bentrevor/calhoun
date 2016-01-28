package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	. "github.com/bentrevor/calhoun/app"
	"github.com/namsral/flag"
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
	rootDir := flag.StringVar(&rootDir, "root-dir", "/home/vagrant/go/src/github.com/bentrevor/calhoun")
	assetPath := flag.StringVar(&assetPath, "asset-path", fmt.Sprintf("%s/assets", rootDir))
	srvPath := flag.StringVar(&srvPath, "srv-path", fmt.Sprintf("%s/images/srv", rootDir))
	flag.Parse()

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/upload_photo", UploadFormHandler)
	http.HandleFunc("/upload", UploadHandler(srvPath))

	// TODO should use a real asset pipeline eventually
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/home/vagrant/go/src/github.com/bentrevor/calhoun/assets"))))

	log.Print("server starting on 8080...\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
