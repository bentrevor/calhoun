package app_test

import (
	"fmt"
	"io"
	"os"

	. "github.com/bentrevor/calhoun/app"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestRoute_ApplyMiddleware(t *testing.T) {
	Describe("Route: ApplyMiddlewareToBase")
	It("applies middleware in order")
	appliedMiddlewares := []string{}

	middlewares := []Middleware{
		func(h CalhounHandler) CalhounHandler {
			fmt.Println("1")
			appliedMiddlewares = append(appliedMiddlewares, "first")
			return h
		},
		func(h CalhounHandler) CalhounHandler {
			fmt.Println("2")
			appliedMiddlewares = append(appliedMiddlewares, "second")
			return h
		},
	}

	route := Route{
		Middlewares: middlewares,
		BaseHandlerFunc: func(_ io.Writer, _ *CalhounRequest) {
			appliedMiddlewares = append(appliedMiddlewares, "base")
		},
	}

	appliedMiddlewares = []string{}
	handler := route.ApplyMiddlewareToBase()
	handler(os.Stdout, &CalhounRequest{})

	// really confused about why this is running the second middleware before the first...
	AssertEquals(t, appliedMiddlewares, []string{"first", "second", "base"})
}
