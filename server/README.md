# Monstera Shop — Backend

Go REST API. Handles auth, products, orders, blog, reviews, and image uploads.

## Stack

- **Go + Gin** — HTTP router
- **pgx/v5** — PostgreSQL driver (connection pool)
- **golang-jwt/jwt/v5** — JWT auth (httpOnly cookie `monstera_token`)
- **golang-migrate** — SQL migrations
- **disintegration/imaging** — Image resize (800×800 JPEG)
- **bcrypt (cost 12)** — Password hashing

## Setup

```bash
# Copy and fill in secrets
cp .env.example .env

# Start PostgreSQL (Docker)
make docker-up

# Apply migrations
make migrate-up

# Start server with hot reload
make dev        # http://localhost:8080
```

### Prerequisites

- Go 1.21+
- Docker
- `air` for hot reload: `go install github.com/air-verse/air@latest`
- `migrate` CLI with postgres tag: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

## Environment

```
# server/.env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/monstera?sslmode=disable
JWT_SECRET=your-secret
ALLOWED_ORIGINS=http://localhost:5173
PORT=8080
```

## Makefile Commands

| Command              | Description                       |
|----------------------|-----------------------------------|
| `make dev`           | Start server with hot reload      |
| `make build`         | Build binary to `bin/server`      |
| `make test`          | Run all tests                     |
| `make docker-up`     | Start PostgreSQL container        |
| `make docker-down`   | Stop PostgreSQL container         |
| `make migrate-up`    | Apply all pending migrations      |
| `make migrate-down`  | Roll back one migration           |

## Module Structure

Each domain lives in `internal/<module>/` with four files:

| File             | Purpose                                      |
|------------------|----------------------------------------------|
| `model.go`       | Structs and request/response types           |
| `repository.go`  | All SQL queries (package-level functions)    |
| `handler.go`     | `Handler` struct; methods are HTTP handlers  |
| `routes.go`      | `Register(...)` wires routes onto gin groups |

## Modules

| Module    | Handles                                        |
|-----------|------------------------------------------------|
| `auth`    | Register, login, logout, me                    |
| `users`   | Update email / password                        |
| `plants`  | Plant catalog + categories (public + admin)    |
| `blog`    | Blog posts (public + admin)                    |
| `orders`  | Checkout, user orders, admin order management  |
| `reviews` | Plant reviews (public + admin)                 |
| `admin`   | Dashboard stats                                |
| `upload`  | Image upload + resize                          |

## Admin User

Promote a user to admin via SQL:

```sql
UPDATE users SET role = 'admin' WHERE email = 'your@email.com';
```
