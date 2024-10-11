package main

import "github.com/yonasketema/blogo/internal/models"

type templateData struct {
	Blog  models.Blog
	Blogs []models.Blog
}
