package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (a *app) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	a.logger.Error(err.Error(), "method", method, "uri", uri, trace, "trace")
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *app) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {

	ts, ok := a.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		a.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)

	if err != nil {
		a.serverError(w, r, err)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (a *app) notFound(w http.ResponseWriter, r *http.Request, data templateData) {
	a.render(w, r, http.StatusNotFound, "notFound.html", data)
}
