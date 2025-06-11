package models

import "time"

type Blog struct {
	Title    string
	Slug     string
	Date     time.Time
	Excerpt  string
	ImageURL string
	Content  string
}
