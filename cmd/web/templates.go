package main

import "github.com/Mike-95/blog_web/pkg/models"

type templateData struct {
	Post  *models.Post
	Posts []*models.Post
}
