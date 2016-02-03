package app

import (
	"io"
	"log"
)

type CalhounServer interface {
	RegisterRoutes()
	Start()
}

type CalhounRenderer interface {
	UploadPhotoForm(io.Writer)
	UploadPhoto(io.Writer)
	ViewPhotos(io.Writer, []Photo)
}

type CalhounHandler func(io.Writer, *CalhounRequest)
type Middleware func(CalhounHandler) CalhounHandler

type CalhounRequest struct {
	Url  string
	Body string
}

type Route struct {
	Path            string
	Middlewares     []Middleware
	BaseHandlerFunc CalhounHandler
}

func LoggingMW(f CalhounHandler) CalhounHandler {
	return func(w io.Writer, r *CalhounRequest) {
		log.Print("\n\nin LoggingMW\n=============\n")
		f(w, r)
	}
}

func LoggingMW2(f CalhounHandler) CalhounHandler {
	return func(w io.Writer, r *CalhounRequest) {
		log.Print("\n\nin LoggingMW2\n=============\n")
		f(w, r)
	}
}

func (route Route) HandlerFunc() CalhounHandler {
	return route.applyMiddleware(0)
}

func (route Route) applyMiddleware(count int) CalhounHandler {
	if count >= len(route.Middlewares) {
		return route.BaseHandlerFunc
	} else {
		return route.Middlewares[count](route.applyMiddleware(count + 1))
	}
}
