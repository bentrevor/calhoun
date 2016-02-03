package web

import (
	"fmt"
	"io"
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
	s.Routes = []Route{
		Route{
			Path:            "/upload",
			BaseHandlerFunc: s.uploadPhoto(),
		},
		Route{
			Path:            "/upload_photo",
			BaseHandlerFunc: s.uploadPhotoForm(),
			Middlewares:     []Middleware{LoggingMW, LoggingMW2},
		},
		Route{
			Path:            "/view_photos",
			BaseHandlerFunc: s.viewPhotos(),
			Middlewares:     []Middleware{LoggingMW},
		},
	}

	s.registerPageRoutes()
	s.registerAssetRoutes()
}

func (WebServer) Start() {
	log.Print("server starting on 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s WebServer) registerPageRoutes() {
	for i := 0; i < len(s.Routes); i++ {
		route := s.Routes[i]
		hf := route.BuildCalhounHandler()

		http.HandleFunc(route.Path, s.buildHttpHandlerFunc(hf))
	}
}

func (s WebServer) buildHttpHandlerFunc(f CalhounHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		calhounReq := CalhounRequest{Url: r.URL.Path, Body: "io.ReadCloser stuff"}
		f(w, &calhounReq)
	}
}

func (s WebServer) registerAssetRoutes() {
	// TODO should use a real asset pipeline eventually
	assetPath := fmt.Sprintf("/%s/", s.AssetPath)
	http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(s.FullAssetPath))))
}

func (s WebServer) uploadPhotoForm() CalhounHandler {
	return func(w io.Writer, _ *CalhounRequest) {
		s.App.UploadPhotoForm(w)
	}
}

func (s WebServer) uploadPhoto() CalhounHandler {
	return func(w io.Writer, r *CalhounRequest) {
		// CalhounRequest.FormFile undefined
		file, _, err := r.FormFile("photoUpload")
		defer file.Close()

		if err != nil {
			fmt.Fprintln(w, "error reading photo upload: ", err)
			return
		}

		s.App.UploadPhoto(w, &file)
	}
}

func (s WebServer) viewPhotos() CalhounHandler {
	return func(w io.Writer, _ *CalhounRequest) {
		s.App.ViewPhotos(w)
	}
}
