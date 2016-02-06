package web_test

import (
	"fmt"
	"regexp"
	"testing"

	. "github.com/bentrevor/calhoun/app"
	. "github.com/bentrevor/calhoun/web"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestWebServer_RegisterRoutes(t *testing.T) {
	assetPath := "fake/assets/path"
	server := WebServer{AssetPath: assetPath}

	Describe("Server: RegisterRoutes")
	It("populates the []Routes on the Server")

	AssertEquals(t, 0, len(server.Routes))
	server.RegisterRoutes()
	Assert(t, 0 < len(server.Routes))

	It("registers an assets route")
	paths := []string{}
	for _, route := range server.Routes {
		paths = append(paths, route.Path)
	}
	includesAssetRoute := anyRegexMatches(paths, assetPath)
	Assert(t, includesAssetRoute)
}

func anyRegexMatches(paths []string, assetPath string) bool {
	for _, path := range paths {
		fmt.Println(path)
		if ok, _ := regexp.MatchString(assetPath, path); ok {
			return true
		}
	}

	return false
}

// really the only thing I want to test for this server is `uploadPhoto` reading a `FormFile`, but
// I'm not sure how to...
