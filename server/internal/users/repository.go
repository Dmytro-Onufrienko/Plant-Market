package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetByEmail(ctx context.Context, pool *pgxpool.Pool, email string) (*User, error) {
	u := &User{}
	err := pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, role, created_at
		 FROM users WHERE email = $1`, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetByID(ctx context.Context, pool *pgxpool.Pool, id string) (*User, error) {
	u := &User{}
	err := pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, role, created_at
		 FROM users WHERE id = $1`, id).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Create(ctx context.Context, pool *pgxpool.Pool, id, email, passwordHash, name string) (*User, error) {
	u := &User{}
	err := pool.QueryRow(ctx,
		`INSERT INTO users (id, email, password_hash, name)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, email, password_hash, name, role, created_at`,
		id, email, passwordHash, name).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateEmail(ctx context.Context, pool *pgxpool.Pool, id, newEmail string) error {
	_, err := pool.Exec(ctx,
		`UPDATE users SET email = $1 WHERE id = $2`, newEmail, id)
	return err
}

func UpdatePassword(ctx context.Context, pool *pgxpool.Pool, id, newHash string) error {
	_, err := pool.Exec(ctx,
		`UPDATE users SET password_hash = $1 WHERE id = $2`, newHash, id)
	return err
}
