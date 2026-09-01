<script setup lang="ts">
import { useRouter } from '#imports';
import { useAuth } from '~/composables/useAuth';
import { parseApiError } from '~/utils/errorUtils';

const router = useRouter();
const auth = useAuth();
const username = ref('');
const email = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

const submit = async () => {
  if (!username.value.trim() || !email.value.trim() || !password.value.trim()) {
    error.value = 'Please fill in all fields.';
    return;
  }

  error.value = '';
  loading.value = true;
  try {
    await auth.register({ username: username.value, email: email.value, password: password.value });
    await router.push('/profile');
  } catch (err: any) {
    error.value = parseApiError(err, 'Registration failed. Please try again.');
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="mx-auto max-w-xl px-4 py-12">
    <div class="rounded-[32px] border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-800 dark:bg-slate-900">
      <h1 class="text-3xl font-black text-slate-900 dark:text-white">Create account</h1>
      <p class="mt-2 text-sm text-slate-500 dark:text-slate-400">Join RecipeJoy to save your favorites and publish recipes.</p>

      <form @submit.prevent="submit" class="mt-8 space-y-5">
        <div>
          <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Username</label>
          <input v-model="username" type="text" placeholder="Username" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
        </div>

        <div>
          <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Email</label>
          <input v-model="email" type="email" placeholder="Email" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
        </div>

        <div>
          <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Password</label>
          <input v-model="password" type="password" placeholder="Password" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
        </div>

        <div v-if="error" class="rounded-2xl bg-red-50 px-4 py-3 text-sm text-red-700 dark:bg-red-900/20 dark:text-red-200">{{ error }}</div>

        <button type="submit" :disabled="loading" class="w-full rounded-full bg-orange-500 px-5 py-3 text-sm font-semibold text-white transition hover:bg-orange-600 disabled:cursor-not-allowed disabled:opacity-60">
          {{ loading ? 'Registering...' : 'Create account' }}
        </button>

        <p class="text-center text-sm text-slate-500 dark:text-slate-400">Already have an account? <NuxtLink to="/login" class="font-semibold text-orange-500">Login</NuxtLink></p>
      </form>
    </div>
  </div>
</template>
