package app

import (
	"io"
	"log"
)

type Middleware func(CalhounHandler) CalhounHandler

func LoggingMW(f CalhounHandler) CalhounHandler {
	return func(w io.Writer, request *CalhounRequest) {
		// log.Printf("\t%s", request.Method, request.Url)
		f(w, request)
	}
}

func LoggingMW2(f CalhounHandler) CalhounHandler {
	return func(w io.Writer, request *CalhounRequest) {
		log.Printf(":2> hit `%s` route", request.Url)
		f(w, request)
	}
}
