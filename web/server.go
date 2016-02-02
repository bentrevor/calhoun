package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bentrevor/calhoun/app"
)

type WebServer struct {
	App           app.CalhounApp
	AssetPath     string
	FullAssetPath string
}

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Middlewares []Middleware
}

type Middleware interface {
	Chain(http.HandlerFunc) http.HandlerFunc
}

type LoggingMW struct{}

func (mw LoggingMW) Chain(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("\n\nhit the logging middleware!!!\n\n")
		f(w, r)
	}
}

type LoggingMW2 struct{}

func (mw LoggingMW2) Chain(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("\n\nhit the logging middleware 2!!!\n\n")
		f(w, r)
	}
}

type LoggingMW3 struct{}

func (mw LoggingMW3) Chain(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("\n\nhit the logging middleware 3!!!\n\n")
		f(w, r)
	}
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
		Route{
			Path:        "/upload_photo",
			HandlerFunc: s.uploadPhotoForm(),
			Middlewares: []Middleware{LoggingMW3{}, LoggingMW2{}},
		},
		Route{
			Path:        "/view_photos",
			HandlerFunc: s.viewPhotos(),
			Middlewares: []Middleware{LoggingMW{}, LoggingMW3{}, LoggingMW3{}, LoggingMW2{}},
		},
		// Route{Path: "/sign_up"},
		// Route{Path: "/login"},
		// Route{Path: "/logout"},
	}

	for i := 0; i < len(routes); i++ {
		route := routes[i]

		http.HandleFunc(route.Path, route.handlerFunc())
	}
}

func (route Route) handlerFunc() http.HandlerFunc {
	if len(route.Middlewares) == 0 {
		return route.HandlerFunc
	} else {
		return route.handlerFuncWithMiddleware(0)
	}
}

func (route Route) handlerFuncWithMiddleware(count int) http.HandlerFunc {
	if count >= len(route.Middlewares) {
		return route.HandlerFunc
	} else {
		return route.Middlewares[count].Chain(route.handlerFuncWithMiddleware(count + 1))
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
