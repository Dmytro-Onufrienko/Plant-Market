package plants

import "time"

type Plant struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Price       int       `json:"price"` // в центах
	Images      []string  `json:"images"`
	CategoryID  *int      `json:"category_id"`
	Stock       int       `json:"stock"`
	Featured    bool      `json:"featured"`
	CreatedAt   time.Time `json:"created_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Filter struct {
	CategorySlug string
	Featured     *bool
	Page         int
	Limit        int
}

type CreateRequest struct {
	Name        string   `json:"name"        binding:"required"`
	Slug        string   `json:"slug"        binding:"required"`
	Description string   `json:"description"`
	Price       int      `json:"price"       binding:"required,min=1"`
	Images      []string `json:"images"`
	CategoryID  *int     `json:"category_id"`
	Stock       int      `json:"stock"       binding:"min=0"`
	Featured    bool     `json:"featured"`
}

type UpdateRequest struct {
	Name        *string  `json:"name"`
	Slug        *string  `json:"slug"`
	Description *string  `json:"description"`
	Price       *int     `json:"price"  binding:"omitempty,min=1"`
	Images      []string `json:"images"`
	CategoryID  *int     `json:"category_id"`
	Stock       *int     `json:"stock"  binding:"omitempty,min=0"`
	Featured    *bool    `json:"featured"`
}
