package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yonasketema/blogo/internal/models"
	"github.com/yonasketema/blogo/internal/validator"
)

type blogCreateForm struct {
	Title   string
	Content string
	// FieldErrors map[string]string
	validator.Validator
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
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		//  🫖️FieldErrors: map[string]string{}
		// 🫖️ Validator: validator.Validator{FieldErrors: },
	}

	//  🫖️ formData.FieldErrors

	formData.CheckFields(validator.NotBlank(formData.Title), "title", "This field cannot be blank")
	formData.CheckFields(validator.MaxChars(formData.Title, 110), "title", "This field cannot be more than 110 characters long")
	formData.CheckFields(validator.NotBlank(formData.Content), "content", "This field cannot be blank")

	if !formData.Valid() {

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
