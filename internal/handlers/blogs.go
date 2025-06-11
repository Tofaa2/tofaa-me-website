package handlers

import (
	"html/template"
	"net/http"
	"time"

	"me.tofaa/internal/models"
)

func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe get blogs from database or a yaml/markdown file?
	blogs := []models.Blog{
		{
			Title:    "Testing!",
			Slug:     "zig-vs-c-experience",
			Date:     time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
			Excerpt:  "My experience with Zig vs C and C++",
			ImageURL: "/static/images/blog/go-programming.jpg",
			Content:  "Full content here...",
		},
	}

	data := struct {
		Title string
		Blogs []models.Blog
	}{
		Title: "Blog Posts",
		Blogs: blogs,
	}

	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/blogs.html",
	))

	tmpl.Execute(w, data)
}
