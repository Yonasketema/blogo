package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/components/nav.html",
		"./ui/html/pages/home.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server errror", http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server errror", http.StatusInternalServerError)
	}

}

func createBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("home ..... "))
}

func viewCreateBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("viewCreateBlog ..... "))
}

func viewBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("blog id %d", id)
	w.Write([]byte(msg))
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /blog/create", createBlog)

	mux.HandleFunc("GET /blog/view/create", viewCreateBlog)
	mux.HandleFunc("GET /blog/view/{id}", viewBlog)

	log.Print("> server running on port :8080 ")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
