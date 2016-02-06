package app_test

import (
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
			return func(w io.Writer, r *CalhounRequest) {
				appliedMiddlewares = append(appliedMiddlewares, "first")
				h(w, r)
			}
		},
		func(h CalhounHandler) CalhounHandler {
			return func(w io.Writer, r *CalhounRequest) {
				appliedMiddlewares = append(appliedMiddlewares, "second")
				h(w, r)
			}
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

	AssertEquals(t, appliedMiddlewares, []string{"first", "second", "base"})
}
