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

func (s CalhounServer) Run(environment string) {
	s.Renderer.RegisterRoutes(s.AssetPath, s.FullAssetPath, s.Store)

	if environment != "test" {
		log.Print("server starting on 8080...\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
