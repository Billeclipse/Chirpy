# Chirpy

Chirpy is a small Go HTTP API built while working through Boot.dev's "Learn HTTP Servers in Go" course. It is a lightweight social backend with users, chirps, JWT authentication, refresh tokens, webhook-driven account upgrades, and a small static app surface.

The project is intentionally compact. Most of the logic lives in [main.go](/home/billy/workspace/Chirpy/main.go), SQL queries are defined under [sql/queries](/home/billy/workspace/Chirpy/sql/queries), and SQLC generates the typed database layer under [internal/database](/home/billy/workspace/Chirpy/internal/database).

## Features

- User registration with hashed passwords
- Login with JWT access tokens and refresh tokens
- Authenticated chirp creation
- Chirp listing, filtering, and sorting
- Single chirp lookup and author-only deletion
- User profile updates
- Chirpy Red upgrades through Polka webhooks
- Admin metrics and reset endpoints
- Static file serving from `/app`

## Stack

- Go
- PostgreSQL
- SQLC
- Goose-style SQL migrations
- JWT (`github.com/golang-jwt/jwt/v5`)
- Argon2 password hashing

## Project Layout

```text
.
├── main.go
├── internal/
│   ├── auth/
│   └── database/
├── sql/
│   ├── queries/
│   └── schema/
├── assets/
├── index.html
└── sqlc.yaml
```

## Environment

Create a local `.env` file with values like:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable
PLATFORM=dev
JWT_SECRET=replace-me
POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e
```

Notes:

- `PLATFORM=dev` enables `POST /admin/reset`
- `POLKA_KEY` is used by the Polka webhook endpoint
- for the Boot.dev lesson flow, webhook behavior depends on whether `POLKA_KEY` is present

## Run

Start PostgreSQL, apply your schema, then run:

```bash
go run .
```

The server listens on:

```text
http://localhost:8080
```

## Test

```bash
go test ./...
```

## API Overview

### Health and Admin

- `GET /api/healthz`
- `GET /admin/metrics`
- `POST /admin/reset`

### Users and Auth

- `POST /api/users`
- `PUT /api/users`
- `POST /api/login`
- `POST /api/refresh`
- `POST /api/revoke`

### Chirps

- `POST /api/chirps`
- `GET /api/chirps`
- `GET /api/chirps/{chirpID}`
- `DELETE /api/chirps/{chirpID}`

`GET /api/chirps` supports:

- `author_id=<uuid>` to filter by author
- `sort=asc|desc` to control `created_at` ordering

### Webhooks

- `POST /api/polka/webhooks`

This endpoint upgrades a user to `is_chirpy_red=true` when Polka sends a `user.upgraded` event.

## Example Requests

Create a user:

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'
```

Log in:

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'
```

List chirps by author in descending order:

```bash
curl "http://localhost:8080/api/chirps?author_id=<user-uuid>&sort=desc"
```

## Notes

This is a course project, so the implementation favors clarity over abstraction. The current code is straightforward to read, easy to extend lesson by lesson, and a good base for adding more endpoints or splitting responsibilities out of `main.go` later.
