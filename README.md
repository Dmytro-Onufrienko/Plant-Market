# Monstera Shop

Online plant store with a SvelteKit frontend and a Go REST API backend.

## Services

| Service  | Stack                      | Port | Directory |
|----------|----------------------------|------|-----------|
| Frontend | SvelteKit + TypeScript     | 5173 | `app/`    |
| Backend  | Go + Gin + PostgreSQL      | 8080 | `server/` |

## Quick Start

```bash
# 1. Start PostgreSQL
cd server && make docker-up

# 2. Run migrations
make migrate-up

# 3. Start Go API (hot reload)
make dev

# 4. Start SvelteKit (new terminal)
cd ../app && npm install && npm run dev
```

Open [http://localhost:5173](http://localhost:5173).

## Project Structure

```
monstera/
├── app/      # SvelteKit frontend
└── server/   # Go backend
```

See `app/README.md` and `server/README.md` for service-specific details.
