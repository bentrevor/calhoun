package app

import (
	"log"
	"net/http"
)

type CalhounServer struct {
	Store    CalhounStore
	Renderer CalhounRenderer

	// need to pass these in separately because the FileServer needs to know where the images
	// are stored.  If it was hard-coded, my app layer would depend on my view layer
	AssetPath     string
	FullAssetPath string
}

func (s CalhounServer) Run() {
	// TODO organize routes
	routes := []string{
		"/upload",
		"/upload_photo",
		"/sign_up",
		"/login",
		"/logout",
		"/view_photos",
		// "/view_photo/:id",
	}

	for i := 0; i < len(routes); i++ {
		http.HandleFunc(routes[i], s.Renderer.Handle(routes[i]))
	}

	// TODO should use a real asset pipeline eventually
	http.Handle(s.AssetPath, http.StripPrefix(s.AssetPath, http.FileServer(http.Dir(s.FullAssetPath))))

	log.Print("server starting on 8080...\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
