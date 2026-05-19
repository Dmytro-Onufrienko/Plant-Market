# Monstera Shop — Frontend

SvelteKit application. Renders server-side, calls the Go API for all data.

## Stack

- **SvelteKit** — full-stack framework (SSR)
- **TypeScript** — all server-side files typed
- **Tailwind CSS** — styling
- **Zod** — form validation
- **Tiptap** — rich text editor (blog)
- **Paraglide** — i18n

## Setup

```bash
npm install
npm run dev        # http://localhost:5173
```

Requires the Go API running at `http://localhost:8080` (see `server/`).

## Environment

```
# app/.env
PUBLIC_API_URL=http://localhost:8080
```

## Scripts

| Command           | Description                     |
|-------------------|---------------------------------|
| `npm run dev`     | Development server (port 5173)  |
| `npm run build`   | Production build                |
| `npm run preview` | Preview production build        |
| `npm run check`   | Type-check with svelte-check    |

## Key Conventions

- All API calls go through `src/lib/api.ts`
- Prices are stored as integer cents; use `formatPrice()` from `src/lib/utils/format.ts`
- Auth is handled via an httpOnly cookie set by the Go API; `hooks.server.ts` calls `GET /api/auth/me` to hydrate `locals.user`
- Load functions transform snake_case API responses to camelCase before passing to components
