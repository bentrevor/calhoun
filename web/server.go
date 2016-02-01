package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bentrevor/calhoun/app"
)

// just take in web request, parse, and pass to Calhoun
type WebServer struct {
	App           app.CalhounApp
	AssetPath     string
	FullAssetPath string
}

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
}

func (s WebServer) RegisterRoutes() {
	s.registerPageRoutes()
	s.registerAssetRoutes()
}

func (s WebServer) Start() {
	log.Print("server starting on 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s WebServer) registerPageRoutes() {
	routes := []Route{
		Route{Path: "/upload", HandlerFunc: s.uploadPhoto()},
		Route{Path: "/upload_photo", HandlerFunc: s.uploadPhotoForm()},
		Route{Path: "/view_photos", HandlerFunc: s.viewPhotos()},
		// Route{Path: "/sign_up"},
		// Route{Path: "/login"},
		// Route{Path: "/logout"},
	}

	for i := 0; i < len(routes); i++ {
		http.HandleFunc(routes[i].Path, routes[i].HandlerFunc)
	}
}

func (s WebServer) registerAssetRoutes() {
	// TODO should use a real asset pipeline eventually
	assetPath := fmt.Sprintf("/%s/", s.AssetPath)
	http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(s.FullAssetPath))))
}

func (s WebServer) uploadPhotoForm() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		s.App.UploadPhotoForm(w)
	}
}

func (s WebServer) uploadPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("photoUpload")
		defer file.Close()

		if err != nil {
			fmt.Fprintln(w, "error reading photo upload: ", err)
			return
		}

		s.App.UploadPhoto(w, &file)
	}
}

func (s WebServer) viewPhotos() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		s.App.ViewPhotos(w)
	}
}
