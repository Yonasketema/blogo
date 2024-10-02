package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/components/nav.html",
		"./ui/html/pages/home.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, r, err)
		http.Error(w, "Internal server errror", http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		a.serverError(w, r, err)
		http.Error(w, "Internal server errror", http.StatusInternalServerError)
	}

}

func (a *app) createBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("home ..... "))
}

func (a *app) viewCreateBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("viewCreateBlog ..... "))
}

func (a *app) viewBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("blog id %d", id)
	w.Write([]byte(msg))
}
