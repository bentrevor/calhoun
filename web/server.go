package web

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
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
			Path: "/upload",
			BaseHandlerFunc: func(w io.Writer, r *CalhounRequest) {
				s.App.UploadPhoto(w, r.UploadFile)
			},
		},
		Route{
			Path: "/upload_photo",
			BaseHandlerFunc: func(w io.Writer, _ *CalhounRequest) {
				s.App.UploadPhotoForm(w)
			},
			Middlewares: []Middleware{LoggingMW},
		},
		Route{
			Path: "/view_photos",
			BaseHandlerFunc: func(w io.Writer, _ *CalhounRequest) {
				s.App.ViewPhotos(w)
			},
			Middlewares: []Middleware{LoggingMW},
		},
	}

	s.registerPageRoutes()
	s.registerAssetRoutes()
}

func (WebServer) Start() {
	log.Print("server starting on 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s WebServer) registerAssetRoutes() {
	// TODO should use a real asset pipeline eventually
	assetPath := fmt.Sprintf("/%s/", s.AssetPath)
	http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(s.FullAssetPath))))
}

func (s WebServer) registerPageRoutes() {
	for i := 0; i < len(s.Routes); i++ {
		route := s.Routes[i]

		calhounHandler := route.ApplyMiddlewareToBase()
		handlerFunc := s.calhounToHttpHandler(calhounHandler, route)

		http.HandleFunc(route.Path, handlerFunc)
	}
}

// ideally, I would be able to define an adapter fn like this for every route, and the Route would
// know how to adapt itself, but then either 1) presentation details (e.g. http) would leak into the
// app, or 2) I would have to add `interface{}` to my types.  This way has some duplication (and a
// lot of indirection...), but I think it best separates app from presentation.
func (s WebServer) calhounToHttpHandler(calhounHandler CalhounHandler, route Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		calhounReq := CalhounRequest{Url: route.Path}

		var file multipart.File
		var err error

		switch route.Path {
		case "/upload":
			file, _, err = r.FormFile("photoUpload")
			defer file.Close()

			if err != nil || file == nil {
				fmt.Fprintln(w, "error reading photo upload: ", err)
				return
			}

			calhounReq.UploadFile = &file
		}

		calhounHandler(w, &calhounReq)
	}
}
