package app

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "some request info:\n%q\n%q\n%q\n%q",
		html.EscapeString(r.URL.Path),
		html.EscapeString(r.Host),
		html.EscapeString(r.RequestURI),
		html.EscapeString(r.RemoteAddr),
	)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body := ""

	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		body = string(bytes)
	}

	fmt.Fprintf(w, "You said: %q", html.EscapeString(body))
}
