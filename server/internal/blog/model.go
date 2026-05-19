package blog

import "time"

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Excerpt    string    `json:"excerpt"`
	CoverImage *string   `json:"cover_image"`
	Published  bool      `json:"published"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateRequest struct {
	Title      string  `json:"title"      binding:"required"`
	Slug       string  `json:"slug"       binding:"required"`
	Content    string  `json:"content"`
	Excerpt    string  `json:"excerpt"`
	CoverImage *string `json:"cover_image"`
	Published  bool    `json:"published"`
}

type UpdateRequest struct {
	Title      *string `json:"title"`
	Slug       *string `json:"slug"`
	Content    *string `json:"content"`
	Excerpt    *string `json:"excerpt"`
	CoverImage *string `json:"cover_image"`
	Published  *bool   `json:"published"`
}
