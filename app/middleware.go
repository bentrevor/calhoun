package app

import (
	"io"
	"log"
)

type Middleware func(CalhounHandler) CalhounHandler

func LoggingMW(f CalhounHandler) CalhounHandler {
	return func(w io.Writer, request *CalhounRequest) {
		log.Printf(":=> hit `%s` route", request.Url)
		f(w, request)
	}
}
