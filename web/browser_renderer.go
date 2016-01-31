package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Yield string
}

type BrowserRenderer struct {
	RespWriter http.ResponseWriter
	ViewsPath  string
}

func (br BrowserRenderer) Handle(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO switch route
		fmt.Fprintf(w, "the route you got was: %s", route)
	}
}

func (br *BrowserRenderer) RenderHtmlFile(filename string) {
	layoutPath := fmt.Sprintf("%s/layout.html", br.ViewsPath)
	filepath := fmt.Sprintf("%s/%s", br.ViewsPath, filename)

	tmpl, err := template.ParseFiles(layoutPath, filepath)

	if err != nil {
		log.Fatal("error on template.ParseFiles: ", err)
	}

	page := Page{Yield: "value of yeld"}
	tmpl.Execute(br.RespWriter, page)
}

func (p Page) GetYield() string {
	return p.Yield
}
