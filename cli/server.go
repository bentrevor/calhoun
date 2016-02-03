package cli

import (
	"io"

	. "github.com/bentrevor/calhoun/app"
)

type CliServer struct {
	App     CalhounApp
	RootDir string
	Routes  []Route
}

type CliHandler func(io.Writer, CliRequest)

func (s CliServer) RegisterRoutes() {
	s.Routes = []Route{
		Route{
			Path:            "-upload",
			BaseHandlerFunc: s.uploadPhoto(),
			Middlewares:     []Middleware{LoggingMW},
		},
	}
}

func (s CliServer) buildHandlerFunc(f CalhoundHandler) CliHandler {
	return func(w io.Writer, r *CliRequest) {
		f(w, r)
	}
}

func (s CliServer) Start() {
	// run command once and quit (for a ReplServer, this could loop)

	// (not sure if I have access to the flag package here...)
	// input := "./calhoun -upload -file /path/to/file"
	// cmd := "-upload"
	// args := "-file /path/to/file"

	// route := s.Routes.Where(Path: cmd)
	// calhounHandler := route.BuildCalhounHandler()
	// handler := s.buildHandlerFunc(calhounHandler)
	// handler(args)
}

func (s CliServer) uploadPhoto() CalhounHandler {
	return func(w io.Writer, r *CalhounRequest) {
		// input == "-file /path/to/file"

		// file := os.Read("/path/to/file")
		// s.App.UploadPhoto(w, &file)
	}
}
