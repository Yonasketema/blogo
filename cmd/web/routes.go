package main

import "net/http"

func (a *app) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /blog", a.createBlog)

	mux.HandleFunc("GET /blog/create", a.viewCreateBlog)
	mux.HandleFunc("GET /blog/{id}", a.viewBlog)

	return mux
}
