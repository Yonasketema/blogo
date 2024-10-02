package main

import (
	"net/http"
	"runtime/debug"
)

func (a *app) serverError(w http.ResponseWriter, r *http.Request, er error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	a.logger.Error(er.Error(), "method", method, "uri", uri, trace, "trace")
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
