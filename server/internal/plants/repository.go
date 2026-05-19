package plants

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAll(ctx context.Context, pool *pgxpool.Pool, f Filter) ([]Plant, int, error) {
	if f.Limit <= 0 {
		f.Limit = 20
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	offset := (f.Page - 1) * f.Limit

	where := "WHERE 1=1"
	args := []any{}
	i := 1

	if f.CategorySlug != "" {
		where += fmt.Sprintf(" AND c.slug = $%d", i)
		args = append(args, f.CategorySlug)
		i++
	}
	if f.Featured != nil {
		where += fmt.Sprintf(" AND p.featured = $%d", i)
		args = append(args, *f.Featured)
		i++
	}

	var total int
	countQ := fmt.Sprintf(
		`SELECT COUNT(*) FROM plants p LEFT JOIN categories c ON p.category_id = c.id %s`, where)
	if err := pool.QueryRow(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	args = append(args, f.Limit, offset)
	q := fmt.Sprintf(`
		SELECT p.id, p.name, p.slug, p.description, p.price, p.images,
		       p.category_id, p.stock, p.featured, p.created_at
		FROM plants p
		LEFT JOIN categories c ON p.category_id = c.id
		%s
		ORDER BY p.created_at DESC
		LIMIT $%d OFFSET $%d`, where, i, i+1)

	rows, err := pool.Query(ctx, q, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	result, err := scanPlants(rows)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func GetBySlug(ctx context.Context, pool *pgxpool.Pool, slug string) (*Plant, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, name, slug, description, price, images,
		       category_id, stock, featured, created_at
		FROM plants WHERE slug = $1`, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanPlants(rows)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func GetByID(ctx context.Context, pool *pgxpool.Pool, id int) (*Plant, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, name, slug, description, price, images,
		       category_id, stock, featured, created_at
		FROM plants WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanPlants(rows)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func Create(ctx context.Context, pool *pgxpool.Pool, req CreateRequest) (*Plant, error) {
	images := req.Images
	if images == nil {
		images = []string{}
	}

	rows, err := pool.Query(ctx, `
		INSERT INTO plants (name, slug, description, price, images, category_id, stock, featured)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, slug, description, price, images, category_id, stock, featured, created_at`,
		req.Name, req.Slug, req.Description, req.Price, images, req.CategoryID, req.Stock, req.Featured)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanPlants(rows)
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func Update(ctx context.Context, pool *pgxpool.Pool, id int, req UpdateRequest) (*Plant, error) {
	existing, err := GetByID(ctx, pool, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Slug != nil {
		existing.Slug = *req.Slug
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.Price != nil {
		existing.Price = *req.Price
	}
	if req.Images != nil {
		existing.Images = req.Images
	}
	if req.CategoryID != nil {
		existing.CategoryID = req.CategoryID
	}
	if req.Stock != nil {
		existing.Stock = *req.Stock
	}
	if req.Featured != nil {
		existing.Featured = *req.Featured
	}

	rows, err := pool.Query(ctx, `
		UPDATE plants
		SET name=$1, slug=$2, description=$3, price=$4, images=$5,
		    category_id=$6, stock=$7, featured=$8
		WHERE id=$9
		RETURNING id, name, slug, description, price, images, category_id, stock, featured, created_at`,
		existing.Name, existing.Slug, existing.Description, existing.Price,
		existing.Images, existing.CategoryID, existing.Stock, existing.Featured, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanPlants(rows)
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func Delete(ctx context.Context, pool *pgxpool.Pool, id int) error {
	tag, err := pool.Exec(ctx, `DELETE FROM plants WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func GetAllCategories(ctx context.Context, pool *pgxpool.Pool) ([]Category, error) {
	rows, err := pool.Query(ctx, `SELECT id, name, slug FROM categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}

func scanPlants(rows pgx.Rows) ([]Plant, error) {
	var result []Plant
	for rows.Next() {
		var p Plant
		var images []string
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.Description, &p.Price, &images,
			&p.CategoryID, &p.Stock, &p.Featured, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		if images == nil {
			images = []string{}
		}
		p.Images = images
		result = append(result, p)
	}
	return result, rows.Err()
}
