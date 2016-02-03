package web

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/bentrevor/calhoun/app"
)

type WebServer struct {
	App           CalhounApp
	AssetPath     string
	FullAssetPath string
	Routes        []Route
}

func (s WebServer) RegisterRoutes() {
	routes := []Route{
		Route{
			Path:            "/upload",
			BaseHandlerFunc: s.uploadPhoto,
		},
		Route{
			Path:            "/upload_photo",
			BaseHandlerFunc: s.uploadPhotoForm,
			Middlewares:     []Middleware{LoggingMW{}, LoggingMW2{}},
		},
		Route{
			Path:            "/view_photos",
			BaseHandlerFunc: s.viewPhotos,
			Middlewares:     []Middleware{LoggingMW{}},
		},
	}

	s.registerPageRoutes(routes)
	s.registerAssetRoutes()
}

func (WebServer) Start() {
	log.Print("server starting on 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (WebServer) registerPageRoutes(routes []Route) {
	for i := 0; i < len(routes); i++ {
		route := routes[i]

		http.HandleFunc(route.Path, route.HandlerFunc())
	}
}

func (s WebServer) registerAssetRoutes() {
	// TODO should use a real asset pipeline eventually
	assetPath := fmt.Sprintf("/%s/", s.AssetPath)
	http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(s.FullAssetPath))))
}

func (s WebServer) uploadPhotoForm(w http.ResponseWriter, _ *http.Request) {
	s.App.UploadPhotoForm(w)
}

func (s WebServer) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("photoUpload")
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, "error reading photo upload: ", err)
		return
	}

	s.App.UploadPhoto(w, &file)
}

func (s WebServer) viewPhotos(w http.ResponseWriter, _ *http.Request) {
	s.App.ViewPhotos(w)
}
