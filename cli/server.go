package cli

import (
	"io"

	"github.com/bentrevor/calhoun/app"
)

type CliServer struct {
	App     app.CalhounApp
	RootDir string
	Routes  []Route
}

type CliHandler func(io.Writer, string)

func (s CliServer) RegisterRoutes() {
	s.Routes = []Route{
		Route{
			Path:            "-upload",
			BaseHandlerFunc: s.uploadPhoto,
			Middlewares:     []app.Middleware{app.LoggingMW{}},
		},
	}
}

func (s CliServer) Start() {
	// run command once and quit (for a ReplServer, this could loop)

	// (not sure if I have access to the flag package here...)
	// input := "./calhoun -upload -file /path/to/file"
	// cmd := "-upload"
	// args := "-file /path/to/file"
	// route := s.Routes.Where(Path: cmd)
	// handler := route.HandlerFunc()
	// handler(args)
}

func (s CliServer) uploadPhoto(w io.Writer, input string) {
	// input == "-file /path/to/file"

	// file := os.Read("/path/to/file")
	// app.UploadPhoto(w, &file)
}
