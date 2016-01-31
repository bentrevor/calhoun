package app_test

import (
	. "github.com/bentrevor/calhoun/app"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

type MockRenderer struct{}

var routesAreRegistered bool

func (r MockRenderer) RegisterRoutes(s, t string, st CalhounStore) {
	routesAreRegistered = true
}

// stupid test, mostly just goofing with mocks
func TestCalhounServer_Run(t *testing.T) {
	Describe("Run")
	It("registers routes")

	renderer := MockRenderer{}
	server := CalhounServer{Renderer: renderer}
	server.Run("test")
	Assert(t, routesAreRegistered)
}
