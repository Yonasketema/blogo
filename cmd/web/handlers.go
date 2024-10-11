package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/yonasketema/blogo/internal/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {

	blogs, err := a.blogs.GetAllBlog()

	if err != nil {
		a.serverError(w, r, err)
		return
	}

	// for _, blog := range blogs {
	// 	fmt.Fprint(w, "%+v\n", blog)
	// }
	//

	data := templateData{
		Blogs: blogs,
	}

	a.render(w, r, http.StatusOK, "home.html", data)

}

func (a *app) createBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	title := "wow go"
	content := "best lang fast and easy"

	id, err := a.blogs.InsertBlog(title, content)
	if err != nil {
		a.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/%d", id), http.StatusSeeOther)
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
	blog, err := a.blogs.GetOneBlog(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			a.serverError(w, r, err)
		}
		return

	}

	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/components/nav.html",
		"./ui/html/pages/blog.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, r, err)
		return
	}

	data := templateData{
		Blog: blog,
	}

	err = ts.ExecuteTemplate(w, "base", data)

	if err != nil {
		a.serverError(w, r, err)
	}

	// fmt.Fprintf(w, "%+v", blog)
	// w.Write([]byte(msg))
}
