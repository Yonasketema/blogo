package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (a *app) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// InfoColor := "\033[1;34m%s %s\033[0m\n"
		// fmt.Printf(InfoColor, r.Host, r.URL.Path)

		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)

		a.logger.Info("request", "ip", ip, "proto", proto, "method", method, "uri", uri)

		next.ServeHTTP(w, r)

	})
}

func (a *app) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				var (
					method = r.Method
					uri    = r.URL.RequestURI()
					trace  = string(debug.Stack())
				)
				err := fmt.Errorf("%s", err)
				a.logger.Error(err.Error(), "method", method, "uri", uri, trace, "trace")

				data := templateData{}
				a.render(w, r, http.StatusInternalServerError, "serverError.html", data)
				// a.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})

}
