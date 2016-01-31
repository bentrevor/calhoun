package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	app "github.com/bentrevor/calhoun/app"
)

type Page struct {
	Yield string
}

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
}

type BrowserRenderer struct {
	ViewsPath string
}

func (br BrowserRenderer) RegisterRoutes(assetPath, fullAssetPath string, store app.CalhounStore) {
	br.registerUserRoutes(store)
	br.registerAssetRoutes(assetPath, fullAssetPath)
}

func (br BrowserRenderer) registerUserRoutes(store app.CalhounStore) {
	routes := []Route{
		Route{
			Path: "/upload",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				file, _, err := r.FormFile("photoUpload")
				defer file.Close()

				if err != nil {
					fmt.Fprintln(w, "error reading photo upload: ", err)
					return
				}

				user := app.User{Id: 1, Name: "God"} // until auth middleware is implemented
				err = store.SavePhoto(user, &file)

				if err == nil {
					br.RenderHtmlFile("upload_success", w)
				} else {
					fmt.Fprintln(w, "error saving photo: ", err)
					return
				}
			}},
		Route{
			Path: "/upload_photo",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				br.RenderHtmlFile("upload_photo_form", w)
			}},
		// Route{Path: "/sign_up"},
		// Route{Path: "/login"},
		// Route{Path: "/logout"},
		// Route{Path: "/view_photos"},
	}

	for i := 0; i < len(routes); i++ {
		http.HandleFunc(routes[i].Path, routes[i].HandlerFunc)
	}
}

func (br BrowserRenderer) registerAssetRoutes(serverAssetPath, fullAssetPath string) {
	// TODO should use a real asset pipeline eventually
	assetPath := fmt.Sprintf("/%s/", serverAssetPath)
	http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(fullAssetPath))))
}

func (br *BrowserRenderer) RenderHtmlFile(filename string, writer http.ResponseWriter) {
	layoutPath := fmt.Sprintf("%s/layout.html", br.ViewsPath)
	filepath := fmt.Sprintf("%s/%s.html", br.ViewsPath, filename)

	tmpl, err := template.ParseFiles(layoutPath, filepath)

	if err != nil {
		log.Fatal("error on template.ParseFiles: ", err)
	}

	tmpl.Execute(writer, Page{})
}
