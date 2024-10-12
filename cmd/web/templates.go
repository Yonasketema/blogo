package main

import (
	"html/template"
	"path/filepath"

	"github.com/yonasketema/blogo/internal/models"
)

type templateData struct {
	Blog  models.Blog
	Blogs []models.Blog
}

func templateCache() (map[string]*template.Template, error) {

	cahce := map[string]*template.Template{}

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

		cahce[name] = ts

	}

	return cahce, nil

}
