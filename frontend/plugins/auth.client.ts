import { defineNuxtPlugin } from '#imports';
import { useAuth } from '~/composables/useAuth';

export default defineNuxtPlugin(async () => {
  const auth = useAuth();
  await auth.initializeAuth();
});
