package web_test

import (
	"net/http"
	"testing"

	. "github.com/bentrevor/calhoun/web"

	. "github.com/bentrevor/calhoun/spec-helper"
)

type FakeMW struct {
	Name string
}

var middlewareCalled = []Middleware{}

func (mw FakeMW) Chain(f http.HandlerFunc) http.HandlerFunc {
	return f
}

// routes
func TestWebServer_Middleware(t *testing.T) {
	// testApp := app.Calhoun{}
	// server := WebServer{App: testApp}

	Describe("Server: middleware")
	It("applies middleware in order before the HandlerFunc")
	AssertEquals(t, 1, 2)
}

// uploadPhoto reads a FormFile

// registers an assets route
