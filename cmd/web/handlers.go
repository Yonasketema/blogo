package main

import (
	"errors"
	"fmt"
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

	data := templateData{}

	if err != nil || id < 1 {
		a.notFound(w, r, data)

		return
	}

	blog, err := a.blogs.GetOneBlog(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w, r, data)
		} else {
			a.serverError(w, r, err)
		}
		return

	}

	data = templateData{
		Blog: blog,
	}
	a.render(w, r, http.StatusOK, "blog.html", data)

	// fmt.Fprintf(w, "%+v", blog)
	// w.Write([]byte(msg))
}
