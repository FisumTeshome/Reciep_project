# RecipeJoy - Food Recipe & Community Platform
Frontend deployed preview: https://reciep-project.vercel.app/
Backend deployed : https://reciep-project.onrender.com
hasura deployed : https://minabpro-hasura.onrender.com

# Recipe Portfolio App

RecipeJoy is a modern full-stack web application designed for food enthusiasts to discover, create, save, rate, and share recipes.

Built with a **Nuxt 3** frontend, **Go (Golang)** REST API backend, **PostgreSQL**, **Hasura GraphQL**, and containerized with **Docker**.

---

## 🌟 Key Features & Recent Updates

### 🔐 Authentication & Session Management
- **User Registration (`/register`)**: Create a new account with a unique username, email, and password.
- **User Login (`/login`)**: Authenticate securely using username/email and password with JWT token response.
- **Persistent & Automatic Token Validation**: Tokens are stored in `localStorage` and validated automatically against `/api/auth/me` on application startup.
- **Global Route Middleware**: Protected routes (`/profile`, `/recipes/create`, `/recipes/:id/edit`) automatically redirect users with missing or expired sessions to `/login`.
- **User Profile Dashboard (`/profile`)**: View user details, published recipes, saved items, and stats.

### 🍳 Recipe Discovery & Management
- **Home & Trending Dishes (`/`)**: Featured pick categories, Ethiopian cuisine highlight, and community creators.
- **Recipe Catalog & Filtering (`/recipes`)**: Browse, search by title/ingredients, and filter by food categories.
- **Recipe Details (`/recipes/:id`)**: Full recipe breakdown including prep/cook times, ingredients list, quick facts, likes, and comment threads.
- **Interactive Ratings & Comments**: Rate recipes (1-5 stars) and join community discussions.
- **Create & Edit Recipes (`/recipes/create`, `/recipes/:id/edit`)**: Interactive forms to create and modify recipes with dynamic ingredient and step lists.
- **Recipe Deletion**: Delete owned recipes with backend synchronization.

### 🛡️ Security & UX Enhancements
- **Conversational Error Messages**: Technical error strings, raw HTTP URLs, ports, methods, and `FetchError` logs are sanitized and replaced with user-friendly messages.
- **Dark Mode Support**: Seamless light/dark theme switching.

---

## 📁 Repository Structure

```text
Reciep_project/
├── frontend/                  ← Nuxt 3 Web Application
│   ├── components/            ← Vue UI Components (Navbar, Footer, RecipeCard)
│   ├── composables/           ← State management & API hooks (useAuth, useRecipes, useTheme)
│   ├── middleware/            ← Global route middleware (auth.global.ts)
│   ├── pages/                 ← Nuxt 3 File-based Routing
│   │   ├── categories.vue
│   │   ├── index.vue
│   │   ├── login.vue
│   │   ├── profile.vue
│   │   ├── register.vue
│   │   └── recipes/
│   │       ├── index.vue
│   │       ├── create.vue
│   │       └── [id]/
│   │           ├── index.vue
│   │           └── edit.vue
│   ├── plugins/               ← Client startup plugins (auth.client.ts)
│   ├── server/                ← Nitro server routes
│   └── utils/                 ← Utility functions (errorUtils.ts)
├── backend/                   ← Go (Golang) REST API
│   ├── controllers/           ← Auth, Recipe, and User controllers
│   ├── database/              ← GORM DB setup & connection
│   ├── models/                ← Data models (User, Recipe, Ingredient, Comment, Rating)
│   ├── routes/                ← Endpoint routing
│   ├── services/              ← Business logic & JWT helper methods
│   └── main.go                ← Application entry point
├── hasura/                    ← Hasura metadata & migrations
├── docker-compose.yml         ← Local development orchestration
└── README.md
```

---

## 🛠️ Tech Stack

- **Frontend**: Nuxt 3, Vue 3, TypeScript, TailwindCSS
- **Backend**: Go (Golang), Net/HTTP, JWT Authentication
- **Database & GraphQL**: PostgreSQL, GORM, Hasura GraphQL Engine
- **DevOps**: Docker, Docker Compose

---

## 🚀 Quick Start (Local Development)

### 1. Run Everything with Docker Compose

```bash
docker compose up --build
```

Once running, access the services:
- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **Go REST API**: [http://localhost:8080](http://localhost:8080)
- **Hasura Console**: [http://localhost:8081](http://localhost:8081)

---

### 2. Running Frontend Separately

```bash
cd frontend
npm install
npm run dev
```

The frontend dev server starts at `http://localhost:3000`.

---

### 3. Running Backend Separately

```bash
cd backend
go mod tidy
go run main.go
```

The Go API server runs at `http://localhost:8080`.

---

## 🔑 Key API Endpoints

### Authentication
- `POST /api/auth/register` — Create a new account
- `POST /api/auth/login` — Log in and obtain JWT token
- `GET /api/auth/me` — Fetch current user profile (requires Bearer token)
- `POST /api/auth/logout` — Invalidate user session

### Recipes
- `GET /api/recipes` — List and search recipes
- `GET /api/recipes/get?id=:id` — Get single recipe detail
- `POST /api/recipes/create` — Create a new recipe (requires auth)
- `PUT /api/recipes/update?id=:id` — Update recipe (requires auth)
- `DELETE /api/recipes/delete?id=:id` — Delete recipe (requires auth)
- `POST /api/recipes/:id/like` — Toggle recipe like (requires auth)
- `POST /api/recipes/:id/bookmark` — Toggle recipe bookmark (requires auth)
- `POST /api/recipes/:id/comments` — Add a comment (requires auth)
- `POST /api/recipes/:id/rating` — Rate a recipe (requires auth)
