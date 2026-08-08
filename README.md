Frontend deployed preview: https://reciep-project.vercel.app/

# Recipe Portfolio App

This repository is organized as a small monorepo for a food recipe platform built with:

- Nuxt 3 frontend
- Go backend
- PostgreSQL + Hasura
- Docker for local development

## Repository structure

```text
portfolio_recipes/
├── frontend/          ← Nuxt 3 app (deployed to Vercel/Netlify)
│   ├── nuxt.config.ts
│   ├── package.json
│   └── ...
├── backend/           ← Go app (deployed to Render)
│   ├── main.go
│   ├── go.mod
│   └── ...
├── hasura/            ← Hasura metadata/migrations
│   ├── migrations/
│   └── metadata/
├── docker-compose.yml ← local development stack
├── .env
├── README.md
└── ...
```

## Why this structure

- Atomic commits: a single PR can update both a Go action handler and the matching Nuxt GraphQL usage.
- Simpler local development: PostgreSQL + Hasura + the Go backend can be started from the root with Docker Compose.
- Clear separation of concerns: frontend, backend, and Hasura metadata are easier to reason about and evolve independently.

## Local development

Start the stack:

```bash
docker compose up --build
```

Then access:

- Frontend: http://localhost:3000
- Hasura Console: http://localhost:8081
- Go API: http://localhost:8080

## Frontend

```bash
cd frontend
npm install
npm run dev
```

## Backend

```bash
cd backend
go mod tidy
go run .
```

## Hasura

Hasura migrations and metadata live under `hasura/` and should be kept in version control so the GraphQL schema and server behavior remain reproducible.

## Notes

This repo is intended to support future feature work and clearer implementation planning. When adding features, keep frontend queries, backend handlers, and Hasura schema changes together so the app stays consistent.
