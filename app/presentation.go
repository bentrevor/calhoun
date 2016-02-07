package app

import (
	"io"
	"mime/multipart"
)

type CalhounServer interface {
	RegisterRoutes()
	Start()
}

// hmmm...
type RenderArgs struct {
	Photos []Photo
}

type CalhounRenderer interface {
	Render(CalhounAction, io.Writer, ...RenderArgs)
}

type CalhounHandler func(io.Writer, *CalhounRequest)

type CalhounRequest struct {
	Url        string
	UploadFile *multipart.File
}

type Route struct {
	Path        string
	Middlewares []Middleware
	Action      CalhounAction
}

func (route Route) ApplyMiddlewareToBase(base CalhounHandler) CalhounHandler {
	return route.applyMiddleware(0, base)
}

func (route Route) applyMiddleware(count int, base CalhounHandler) CalhounHandler {
	if count >= len(route.Middlewares) {
		return base
	} else {
		return route.Middlewares[count](route.applyMiddleware(count+1, base))
	}
}
