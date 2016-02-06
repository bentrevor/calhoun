package app

import (
	"io"
	"mime/multipart"
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

type CalhounRequest struct {
	Url string
	// Body     string
	UploadFile *multipart.File
}

type Route struct {
	Path            string
	Middlewares     []Middleware
	BaseHandlerFunc CalhounHandler
}

func (route Route) ApplyMiddlewareToBase() CalhounHandler {
	return route.applyMiddleware(0)
}

func (route Route) applyMiddleware(count int) CalhounHandler {
	if count >= len(route.Middlewares) {
		return route.BaseHandlerFunc
	} else {
		return route.Middlewares[count](route.applyMiddleware(count + 1))
	}
}
