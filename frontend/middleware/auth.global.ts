import { defineNuxtRouteMiddleware, navigateTo } from '#imports';
import { useAuth } from '~/composables/useAuth';

export default defineNuxtRouteMiddleware(async (to) => {
  const auth = useAuth();
  if (process.client) {
    await auth.initializeAuth();
  }

  const protectedRoutes = ['/profile', '/recipes/create'];
  const isEditRoute = to.path.startsWith('/recipes/') && to.path.endsWith('/edit');
  const isProtectedRoute = protectedRoutes.includes(to.path) || isEditRoute;

  const guestOnlyRoutes = ['/login', '/register'];
  const isGuestOnlyRoute = guestOnlyRoutes.includes(to.path);

  if (isProtectedRoute && !auth.isLoggedIn.value) {
    return navigateTo('/login');
  }

  if (isGuestOnlyRoute && auth.isLoggedIn.value) {
    return navigateTo('/profile');
  }
});
