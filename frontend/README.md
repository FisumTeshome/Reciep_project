# RecipeJoy - Frontend (Nuxt 3)

The frontend application for RecipeJoy built with Nuxt 3, TypeScript, Vue 3, and TailwindCSS.

## Features Included

- **User Authentication**: Login (`/login`) & Registration (`/register`) forms with user-friendly error sanitization.
- **Session Validation**: Automatic token checking on app startup (`plugins/auth.client.ts`) and global route protection (`middleware/auth.global.ts`).
- **Profile Dashboard**: User profile dashboard (`/profile`) with user stats, recipes, and logout.
- **Recipe Management**: Explore recipes (`/recipes`), view recipe details (`/recipes/:id`), publish new recipes (`/recipes/create`), and edit existing recipes (`/recipes/:id/edit`).
- **Sanitized Errors**: Utility function `parseApiError` strips raw technical FetchErrors, HTTP methods, and URL endpoints.

## Development Setup

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

Visit `http://localhost:3000` in your browser.

## Build & Production

```bash
# Build production bundle
npm run build

# Preview production build locally
npm run preview
```
