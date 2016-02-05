package cli

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	. "github.com/bentrevor/calhoun/app"
)

type ConsoleServer struct {
	App    CalhounApp
	Args   []string
	Routes []Route
}

type ConsoleHandler func(io.Writer, ConsoleRequest)
type ConsoleRequest struct {
	Url string
}

func (s ConsoleServer) RegisterRoutes() {
	s.Routes = []Route{
		Route{
			Path: "upload",
			BaseHandlerFunc: func(w io.Writer, r *CalhounRequest) {
				s.App.UploadPhoto(w, r.UploadFile)
			},
			Middlewares: []Middleware{LoggingMW2},
		},
	}
}

func (s ConsoleServer) Start() {
	// here is where I would specify the console interface if I actually planned on using the
	// from the command line.  For now it will just expect a command to look like `./calhoun
	// -ui=cli upload file=/path/to/file`
	url := s.Args[0]
	filepath := strings.SplitAfter(s.Args[1], "=")[1] // file=/path/to/file, so I need everything after the =

	var file multipart.File

	switch url {
	case "upload":
		file, err := os.Open(filepath)
		defer file.Close()

		if err != nil {
			log.Fatal("error reading photo upload: ", err)
		}
	}

	route := s.routeWithPath(url)
	calhounHandler := route.ApplyMiddlewareToBase()
	request := CalhounRequest{UploadFile: &file}
	calhounHandler(os.Stdout, &request)
}

func (s ConsoleServer) routeWithPath(url string) Route {
	for i := 0; i < len(s.Routes); i++ {
		route := s.Routes[i]

		if route.Path == url {
			return route
		}
	}

	return Route{}
}
