package reviews

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetByPlant(ctx context.Context, pool *pgxpool.Pool, plantID int) ([]ReviewWithUser, error) {
	rows, err := pool.Query(ctx, `
		SELECT r.id, r.user_id, r.plant_id, r.rating, r.text, r.created_at, u.name
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		WHERE r.plant_id = $1
		ORDER BY r.created_at DESC`, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []ReviewWithUser
	for rows.Next() {
		var r ReviewWithUser
		if err := rows.Scan(&r.ID, &r.UserID, &r.PlantID, &r.Rating, &r.Text, &r.CreatedAt, &r.UserName); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, rows.Err()
}

func Create(ctx context.Context, pool *pgxpool.Pool, userID string, plantID, rating int, text string) error {
	_, err := pool.Exec(ctx,
		`INSERT INTO reviews (user_id, plant_id, rating, text) VALUES ($1, $2, $3, $4)`,
		userID, plantID, rating, text)
	return err
}

func GetAll(ctx context.Context, pool *pgxpool.Pool) ([]ReviewWithDetails, error) {
	rows, err := pool.Query(ctx, `
		SELECT r.id, r.user_id, r.plant_id, r.rating, r.text, r.created_at,
		       u.name, u.email, p.name, p.slug
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		JOIN plants p ON r.plant_id = p.id
		ORDER BY r.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanWithDetails(rows)
}

func GetLatest(ctx context.Context, pool *pgxpool.Pool, limit int) ([]ReviewWithDetails, error) {
	rows, err := pool.Query(ctx, `
		SELECT r.id, r.user_id, r.plant_id, r.rating, r.text, r.created_at,
		       u.name, u.email, p.name, p.slug
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		JOIN plants p ON r.plant_id = p.id
		ORDER BY r.created_at DESC
		LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews, err := scanWithDetails(rows)
	if err != nil {
		return nil, err
	}
	if reviews == nil {
		reviews = []ReviewWithDetails{}
	}
	return reviews, nil
}

func Delete(ctx context.Context, pool *pgxpool.Pool, id int) error {
	tag, err := pool.Exec(ctx, `DELETE FROM reviews WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func CheckExists(ctx context.Context, pool *pgxpool.Pool, userID string, plantID int) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM reviews WHERE user_id=$1 AND plant_id=$2)`,
		userID, plantID).Scan(&exists)
	return exists, err
}

func scanWithDetails(rows pgx.Rows) ([]ReviewWithDetails, error) {
	var reviews []ReviewWithDetails
	for rows.Next() {
		var r ReviewWithDetails
		if err := rows.Scan(&r.ID, &r.UserID, &r.PlantID, &r.Rating, &r.Text, &r.CreatedAt,
			&r.UserName, &r.UserEmail, &r.PlantName, &r.PlantSlug); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, rows.Err()
}
