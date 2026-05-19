# Monstera Shop — Project Guide

## Architecture

Two separate services:

| Service | Stack | Port | Directory |
|---|---|---|---|
| Frontend | SvelteKit + TypeScript | 5173 | `app/` |
| Backend  | Go + Gin + pgx         | 8080 | `server/` |

Communication: SvelteKit server-side load functions and form actions call the Go REST API. No direct DB access from the frontend.

---

## Local Setup

```bash
# 1. Start PostgreSQL (Docker)
cd server && make docker-up

# 2. Apply migrations
make migrate-up

# 3. Start Go server (hot reload via air)
make dev          # :8080

# 4. Start SvelteKit (separate terminal)
cd app && npm run dev   # :5173
```

---

## Tech Stack

### Backend (`server/`)
- **Go + Gin** — HTTP router
- **pgx/v5** — PostgreSQL driver (connection pool)
- **golang-jwt/jwt/v5** — JWT auth (httpOnly cookie `monstera_token`)
- **golang-migrate** — SQL migrations
- **disintegration/imaging** — Image resize (800×800 JPEG)
- **bcrypt (cost 12)** — Password hashing
- **godotenv** — `.env` loading

### Frontend (`app/`)
- **SvelteKit** — Full-stack framework (server-side rendering)
- **TypeScript** — All server files typed
- **Zod** — Form validation (client-side, mirrors Go validation)
- **Tiptap** — Rich text editor (blog content)
- **i18n** — Internationalisation via `$lib/i18n`

---

## Project Structure

```
monstera/
├── app/                         # SvelteKit frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api.ts           # Go API client (all fetch calls)
│   │   │   ├── utils/format.ts  # formatPrice, parseImages
│   │   │   ├── stores/cart.ts   # Cart state (localStorage)
│   │   │   └── server/
│   │   │       └── validations/ # Zod schemas (form UX only)
│   │   ├── hooks.server.ts      # Auth: calls GET /api/auth/me
│   │   ├── app.d.ts             # App.Locals + API types
│   │   └── routes/              # SvelteKit pages + server files
│   └── .env                     # PUBLIC_API_URL=http://localhost:8080
│
└── server/                      # Go backend
    ├── cmd/server/main.go        # Entry point
    ├── internal/
    │   ├── config/config.go      # Env loading
    │   ├── db/db.go              # pgx pool init only
    │   │   └── migrations/       # 001_init.up/down.sql
    │   ├── middleware/           # CORS, Logger, Auth, AdminRequired
    │   ├── services/             # JWT, Password (bcrypt)
    │   ├── httputil/respond.go   # Shared HTTP response helpers + IntQuery
    │   ├── router/router.go      # Thin assembler — calls each module's Register()
    │   │
    │   ├── auth/                 # Module: auth
    │   │   ├── model.go          # RegisterRequest, LoginRequest
    │   │   ├── handler.go        # Register, Login, Logout, Me
    │   │   └── routes.go
    │   ├── users/                # Module: users
    │   │   ├── model.go          # User, UserPublic, UpdateEmailRequest, UpdatePasswordRequest
    │   │   ├── repository.go     # GetByEmail, GetByID, Create, UpdateEmail, UpdatePassword
    │   │   ├── handler.go        # UpdateEmail, UpdatePassword
    │   │   └── routes.go
    │   ├── plants/               # Module: plants + categories
    │   │   ├── model.go          # Plant, Category, Filter, CreateRequest, UpdateRequest
    │   │   ├── repository.go     # GetAll, GetByID, GetBySlug, Create, Update, Delete, GetAllCategories
    │   │   ├── handler.go        # List, GetByID, GetBySlug, ListCategories + Admin CRUD
    │   │   └── routes.go
    │   ├── blog/                 # Module: blog
    │   │   ├── model.go          # Post, CreateRequest, UpdateRequest
    │   │   ├── repository.go     # GetPublished, GetAll, GetPublishedBySlug, GetByID, Create, Update, Delete
    │   │   ├── handler.go        # List, GetBySlug + Admin CRUD
    │   │   └── routes.go
    │   ├── orders/               # Module: orders
    │   │   ├── model.go          # Order, OrderItem, ShippingAddress, CheckoutRequest, UpdateStatusRequest, …
    │   │   ├── repository.go     # Create, GetByUser, GetByID, GetAll, GetRecent, UpdateStatus
    │   │   ├── handler.go        # Checkout, ListMine, GetMine + Admin list/status
    │   │   └── routes.go
    │   ├── reviews/              # Module: reviews
    │   │   ├── model.go          # Review, ReviewWithUser, ReviewWithDetails, CreateReviewRequest
    │   │   ├── repository.go     # GetByPlant, Create, GetAll, GetLatest, Delete, CheckExists
    │   │   ├── handler.go        # Latest, Create + Admin list/delete
    │   │   └── routes.go
    │   ├── admin/                # Module: admin stats
    │   │   ├── handler.go        # Stats (aggregates counts + recent orders)
    │   │   └── routes.go
    │   └── upload/               # Module: image upload
    │       ├── handler.go        # Upload (resize → 800×800 JPEG)
    │       └── routes.go
    ├── static/uploads/plants/    # Uploaded images (served at /uploads/*)
    ├── .env                      # Local secrets (not in git)
    ├── .env.example              # Keys only (in git)
    ├── .air.toml                 # Air hot reload config
    ├── Makefile                  # dev, docker-up/down, migrate, build, test
    └── docker/docker-compose.yml # PostgreSQL 16 Alpine
```

### Module conventions

Each module owns its models, repository (SQL), handler, and route registration:

- `model.go` — structs + request types for this domain only
- `repository.go` — all SQL (package-level functions, no interface layer)
- `handler.go` — `Handler` struct with `*pgxpool.Pool`; methods are HTTP handlers
- `routes.go` — `Register(...)` wires handler methods onto gin groups

Cross-module imports are allowed one-directionally (e.g. `plants` imports `reviews` for `GetByPlant`; `admin` imports `orders` for `GetRecent`). Circular imports are not allowed.

---

## REST API Summary

```
# Public
GET  /api/plants                  ?category=&page=&featured=&limit=
GET  /api/plants/by-id/:id
GET  /api/plants/:slug
GET  /api/categories
GET  /api/blog
GET  /api/blog/:slug
GET  /api/reviews/latest          ?limit=

# Auth
POST /api/auth/register
POST /api/auth/login
POST /api/auth/logout
GET  /api/auth/me                 (requires cookie)

# User (auth required)
GET  /api/orders
GET  /api/orders/:id
POST /api/orders
PUT  /api/users/email
PUT  /api/users/password
POST /api/reviews

# Admin (role=admin required)
GET  /api/admin/stats
POST /api/admin/plants
PUT  /api/admin/plants/:id
DELETE /api/admin/plants/:id
POST /api/admin/blog
PUT  /api/admin/blog/:id
DELETE /api/admin/blog/:id
GET  /api/admin/orders
PUT  /api/admin/orders/:id/status
GET  /api/admin/reviews
DELETE /api/admin/reviews/:id
POST /api/admin/upload
```

---

## Key Decisions

### Auth
- JWT stored in **httpOnly cookie** `monstera_token` (set by Go)
- SvelteKit **forwards** the cookie to Go on every server-side request
- `hooks.server.ts` validates auth via `GET /api/auth/me` (no DB in frontend)
- Login/register: Go sets `Set-Cookie`, SvelteKit re-sets it via `cookies.set()` for same-origin access

### Data Flow
- All DB access is in Go — SvelteKit has **no database**
- Load functions call Go API and **transform** snake_case → camelCase before passing to components
- Form actions translate form data → JSON → Go API

### Images
- Uploaded via `POST /api/admin/upload` (multipart) → Go resizes to 800×800 JPEG
- Stored in `server/static/uploads/plants/`
- Served by Go at `/uploads/*`
- `parseImages()` handles both `string[]` (from Go) and JSON string (legacy)

### Prices
- Always stored and transferred as **integer cents** (€10.00 = `1000`)
- `formatPrice(cents)` renders as `€ 10.00`

### Shipping Address Fields
- Go API: `{ full_name, street, city, zip, country }`
- Frontend components: `{ name, address, city, country, postalCode }`
- Translation happens in `orders/[id]/+page.server.ts` (load) and `checkout/+page.server.ts` (action)

---

## Gotchas

- **CORS + credentials**: Go CORS middleware checks exact origin match (from `ALLOWED_ORIGINS`). Must be `http://localhost:5173` in dev.
- **sslmode=disable**: Required in `DATABASE_URL` for local Docker PostgreSQL.
- **Cookie domain**: In dev, cookies are set for `localhost` — both services must run on `localhost` (not `127.0.0.1`).
- **`migrate` CLI**: Must be installed with postgres tag: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- **Admin user**: Set via SQL: `UPDATE users SET role='admin' WHERE email='...'`
