package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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

	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	fieldErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 300 {
		fieldErrors["title"] = "This field cannot be more than 300 characters long"
	}

	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field cannot be blank"
	}
	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	id, err := a.blogs.InsertBlog(title, content)
	if err != nil {
		a.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/%d", id), http.StatusSeeOther)
}

func (a *app) viewCreateBlog(w http.ResponseWriter, r *http.Request) {

	data := templateData{}
	a.render(w, r, http.StatusOK, "create.html", data)

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
