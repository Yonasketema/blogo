package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home ..... "))
}
func createBlog(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /blog/create", createBlog)

	mux.HandleFunc("GET /blog/view/create", viewCreateBlog)
	mux.HandleFunc("GET /blog/view/{id}", viewBlog)

	log.Print("> server running on port :8080 ")

	err := http.ListenAndServe(":400", mux)

	log.Fatal(err)
}
