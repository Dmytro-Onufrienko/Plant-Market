package blog

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPublished(ctx context.Context, pool *pgxpool.Pool) ([]Post, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, title, slug, content, excerpt, cover_image, published, created_at
		FROM blog_posts
		WHERE published = true
		ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func GetAll(ctx context.Context, pool *pgxpool.Pool) ([]Post, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, title, slug, content, excerpt, cover_image, published, created_at
		FROM blog_posts
		ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func GetPublishedBySlug(ctx context.Context, pool *pgxpool.Pool, slug string) (*Post, error) {
	return getBySlug(ctx, pool, slug, true)
}

func GetByID(ctx context.Context, pool *pgxpool.Pool, id int) (*Post, error) {
	p := &Post{}
	err := pool.QueryRow(ctx, `
		SELECT id, title, slug, content, excerpt, cover_image, published, created_at
		FROM blog_posts WHERE id = $1`, id).
		Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Excerpt, &p.CoverImage, &p.Published, &p.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return p, err
}

func Create(ctx context.Context, pool *pgxpool.Pool, req CreateRequest) (*Post, error) {
	p := &Post{}
	err := pool.QueryRow(ctx, `
		INSERT INTO blog_posts (title, slug, content, excerpt, cover_image, published)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, slug, content, excerpt, cover_image, published, created_at`,
		req.Title, req.Slug, req.Content, req.Excerpt, req.CoverImage, req.Published).
		Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Excerpt, &p.CoverImage, &p.Published, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func Update(ctx context.Context, pool *pgxpool.Pool, id int, req UpdateRequest) (*Post, error) {
	existing, err := GetByID(ctx, pool, id)
	if err != nil || existing == nil {
		return nil, err
	}

	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Slug != nil {
		existing.Slug = *req.Slug
	}
	if req.Content != nil {
		existing.Content = *req.Content
	}
	if req.Excerpt != nil {
		existing.Excerpt = *req.Excerpt
	}
	if req.CoverImage != nil {
		existing.CoverImage = req.CoverImage
	}
	if req.Published != nil {
		existing.Published = *req.Published
	}

	p := &Post{}
	err = pool.QueryRow(ctx, `
		UPDATE blog_posts
		SET title=$1, slug=$2, content=$3, excerpt=$4, cover_image=$5, published=$6
		WHERE id=$7
		RETURNING id, title, slug, content, excerpt, cover_image, published, created_at`,
		existing.Title, existing.Slug, existing.Content, existing.Excerpt,
		existing.CoverImage, existing.Published, id).
		Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Excerpt, &p.CoverImage, &p.Published, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func Delete(ctx context.Context, pool *pgxpool.Pool, id int) error {
	tag, err := pool.Exec(ctx, `DELETE FROM blog_posts WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func getBySlug(ctx context.Context, pool *pgxpool.Pool, slug string, onlyPublished bool) (*Post, error) {
	q := `SELECT id, title, slug, content, excerpt, cover_image, published, created_at
	      FROM blog_posts WHERE slug = $1`
	if onlyPublished {
		q += ` AND published = true`
	}
	p := &Post{}
	err := pool.QueryRow(ctx, q, slug).
		Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Excerpt, &p.CoverImage, &p.Published, &p.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return p, err
}

func scanPosts(rows pgx.Rows) ([]Post, error) {
	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Excerpt,
			&p.CoverImage, &p.Published, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}
