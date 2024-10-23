package main

import (
	"html/template"
	"path/filepath"

	"github.com/yonasketema/blogo/internal/models"
)

type templateData struct {
	Blog  models.Blog
	Blogs []models.Blog
	Form  any
}

func templateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/pages/base.html",
			"./ui/html/components/nav.html",
			page,
		}

		ts, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil

}
