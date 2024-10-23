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

type blogCreateForm struct {
	Title       string
	Content     string
	FieldErrors map[string]string
}

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

	formData := blogCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(formData.Title) == "" {
		formData.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(formData.Title) > 110 {
		formData.FieldErrors["title"] = "This field cannot be more than 110 characters long"
	}

	if strings.TrimSpace(formData.Content) == "" {
		formData.FieldErrors["content"] = "This field cannot be blank"
	}

	if len(formData.FieldErrors) > 0 {

		data := templateData{}
		data.Form = formData
		a.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		fmt.Fprint(w, formData.FieldErrors)
		return
	}

	id, err := a.blogs.InsertBlog(formData.Title, formData.Content)
	if err != nil {
		a.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/%d", id), http.StatusSeeOther)
}

func (a *app) viewCreateBlog(w http.ResponseWriter, r *http.Request) {

	data := templateData{}
	data.Form = blogCreateForm{}
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
