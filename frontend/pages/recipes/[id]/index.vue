<script setup lang="ts">
import { useRoute, useRouter } from '#imports';
import { useRecipes } from '~/composables/useRecipes';
import { useAuth } from '~/composables/useAuth';
import { useRuntimeConfig } from '#imports';
import { parseApiError } from '~/utils/errorUtils';

const route = useRoute();
const router = useRouter();
const auth = useAuth();
const { recipe, loading, error, fetchRecipe } = useRecipes();
const config = useRuntimeConfig();
const apiBase = config.public.apiBase;
const recipeId = route.params.id as string;

const deleteLoading = ref(false);
const deleteError = ref('');
const isLiked = ref(false);
const isBookmarked = ref(false);
const likeLoading = ref(false);
const bookmarkLoading = ref(false);
const comments = ref<any[]>([]);
const newComment = ref('');
const commentLoading = ref(false);
const userRating = ref(0);
const ratingLoading = ref(false);

if (recipeId) {
  fetchRecipe(recipeId).catch(() => null);
}

const toggleLike = async () => {
  if (!auth.isLoggedIn) {
    await router.push('/login');
    return;
  }

  likeLoading.value = true;
  try {
    const method = isLiked.value ? 'DELETE' : 'POST';
    const { error: fetchError } = await useFetch(`${apiBase}/api/recipes/${recipeId}/like`, {
      method,
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
      },
    });

    if (!fetchError.value) {
      isLiked.value = !isLiked.value;
      if (recipe.value) {
        recipe.value.likes = isLiked.value ? (recipe.value.likes || 0) + 1 : Math.max(0, (recipe.value.likes || 0) - 1);
      }
    }
  } catch (err) {
    console.error('Failed to toggle like:', err);
  } finally {
    likeLoading.value = false;
  }
};

const toggleBookmark = async () => {
  if (!auth.isLoggedIn) {
    await router.push('/login');
    return;
  }

  bookmarkLoading.value = true;
  try {
    const method = isBookmarked.value ? 'DELETE' : 'POST';
    const { error: fetchError } = await useFetch(`${apiBase}/api/recipes/${recipeId}/bookmark`, {
      method,
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
      },
    });

    if (!fetchError.value) {
      isBookmarked.value = !isBookmarked.value;
    }
  } catch (err) {
    console.error('Failed to toggle bookmark:', err);
  } finally {
    bookmarkLoading.value = false;
  }
};

const deleteRecipe = async () => {
  if (!confirm('Are you sure you want to delete this recipe?')) {
    return;
  }

  deleteLoading.value = true;
  deleteError.value = '';

  try {
    const { error: fetchError } = await useFetch(`${apiBase}/api/recipes/delete?id=${encodeURIComponent(recipeId)}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
      },
    });

    if (fetchError.value) {
      throw fetchError.value;
    }

    await router.push('/recipes');
  } catch (err: any) {
    deleteError.value = parseApiError(err, 'Failed to delete recipe.');
  } finally {
    deleteLoading.value = false;
  }
};

const addComment = async () => {
  if (!auth.isLoggedIn) {
    await router.push('/login');
    return;
  }

  if (!newComment.value.trim()) {
    return;
  }

  commentLoading.value = true;
  try {
    const { data, error: fetchError } = await useFetch(`${apiBase}/api/recipes/${recipeId}/comments`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
        'Content-Type': 'application/json',
      },
      body: {
        content: newComment.value,
      },
    });

    if (!fetchError.value && data.value) {
      comments.value.unshift(data.value);
      newComment.value = '';
    }
  } catch (err) {
    console.error('Failed to add comment:', err);
  } finally {
    commentLoading.value = false;
  }
};

const setRating = async (rating: number) => {
  if (!auth.isLoggedIn) {
    await router.push('/login');
    return;
  }

  ratingLoading.value = true;
  try {
    const { data, error: fetchError } = await useFetch(`${apiBase}/api/recipes/${recipeId}/rating`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
        'Content-Type': 'application/json',
      },
      body: {
        rating,
      },
    });

    if (!fetchError.value && data.value) {
      userRating.value = rating;
      if (recipe.value) {
        recipe.value.rating = data.value.average_rating || rating;
      }
    }
  } catch (err) {
    console.error('Failed to set rating:', err);
  } finally {
    ratingLoading.value = false;
  }
};
</script>

<template>
  <div class="space-y-8 px-4 py-12 sm:px-6 lg:px-8">
    <section v-if="loading" class="rounded-[32px] border border-slate-200 bg-white p-8 shadow-sm dark:border-slate-800 dark:bg-slate-900">
      Loading recipe...
    </section>

    <section v-else-if="error" class="rounded-[32px] border border-red-200 bg-red-50 p-8 text-red-700 shadow-sm dark:border-red-900 dark:bg-red-900/20">
      {{ error }}
    </section>

    <div v-if="deleteError" class="rounded-2xl bg-red-50 px-4 py-3 text-sm text-red-700 dark:bg-red-900/20 dark:text-red-200">
      {{ deleteError }}
    </div>

    <section v-else-if="recipe" class="overflow-hidden rounded-[32px] border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <img :src="recipe.featured_image || 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=1600&q=80'" :alt="recipe.title" class="h-[420px] w-full object-cover" />
      <div class="p-6 md:p-8">
        <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <div class="mb-3 inline-flex rounded-full bg-orange-100 px-3 py-1 text-xs font-bold uppercase tracking-[0.18em] text-orange-600 dark:bg-orange-500/10 dark:text-orange-300">
              {{ recipe.category?.name || recipe.category || 'Uncategorized' }}
            </div>
            <h1 class="text-4xl font-black text-slate-900 dark:text-white">{{ recipe.title }}</h1>
          </div>

          <div class="flex gap-3">
            <button @click="toggleLike" :disabled="likeLoading" class="rounded-full border border-slate-200 px-4 py-2 text-sm font-semibold text-slate-700 dark:border-slate-700 dark:text-slate-200 disabled:cursor-not-allowed disabled:opacity-60">
              {{ isLiked ? '❤️ Liked' : '🤍 Like' }}
            </button>
            <button @click="toggleBookmark" :disabled="bookmarkLoading" class="rounded-full border border-slate-200 px-4 py-2 text-sm font-semibold text-slate-700 dark:border-slate-700 dark:text-slate-200 disabled:cursor-not-allowed disabled:opacity-60">
              {{ isBookmarked ? '🔖 Saved' : '📑 Save' }}
            </button>
            <button class="rounded-full bg-orange-500 px-4 py-2 text-sm font-semibold text-white">👨‍🍳 Start Cooking</button>
            <template v-if="auth.isLoggedIn && recipe?.user?.username === auth.user?.username">
              <NuxtLink :to="`/recipes/${recipeId}/edit`" class="rounded-full border border-slate-200 px-4 py-2 text-sm font-semibold text-slate-700 dark:border-slate-700 dark:text-slate-200">
                ✏️ Edit
              </NuxtLink>
              <button @click="deleteRecipe" :disabled="deleteLoading" class="rounded-full border border-red-300 bg-red-50 px-4 py-2 text-sm font-semibold text-red-600 dark:border-red-700 dark:bg-red-900/20 dark:text-red-400 disabled:cursor-not-allowed disabled:opacity-60">
                {{ deleteLoading ? 'Deleting...' : '🗑️ Delete' }}
              </button>
            </template>
          </div>
        </div>

        <div class="mt-6 flex flex-wrap items-center gap-4 text-sm text-slate-600 dark:text-slate-300">
          <span>👨‍🍳 {{ recipe.user?.username || recipe.creator || 'Guest' }}</span>
          <span>⏱ {{ recipe.prep_time || recipe.prepTime || 0 }} min prep</span>
          <span>🍳 {{ recipe.cook_time || recipe.cookTime || 0 }} min cook</span>
          <span>⭐ {{ recipe.rating || 0 }} / 5</span>
          <span>❤️ {{ recipe.likes || recipe.like_count || 0 }} likes</span>
        </div>

        <!-- Rating Section -->
        <div class="mt-6 rounded-2xl bg-slate-50 p-4 dark:bg-slate-800">
          <div class="flex items-center gap-3">
            <span class="text-sm font-semibold text-slate-700 dark:text-slate-200">Rate this recipe:</span>
            <div class="flex gap-1">
              <button 
                v-for="star in 5" 
                :key="star"
                @click="setRating(star)"
                :disabled="ratingLoading"
                class="text-2xl transition hover:scale-110 disabled:cursor-not-allowed disabled:opacity-60"
              >
                {{ star <= userRating ? '⭐' : '☆' }}
              </button>
            </div>
          </div>
          <p v-if="userRating > 0" class="mt-2 text-xs text-slate-500 dark:text-slate-400">
            Your rating: {{ userRating }}/5
          </p>
        </div>

        <p class="mt-6 max-w-3xl text-lg text-slate-600 dark:text-slate-300">
          {{ recipe.description }}
        </p>
      </div>
    </section>

    <section v-if="recipe" class="grid gap-8 lg:grid-cols-[1fr_0.8fr]">
      <div class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-2xl font-black text-slate-900 dark:text-white">Ingredients</h2>
        <ul class="mt-5 space-y-3">
          <li v-for="ingredient in recipe.recipe_ingredients || recipe.ingredients || []" :key="ingredient.id || ingredient.name" class="flex items-center gap-3 rounded-2xl bg-slate-50 px-4 py-3 dark:bg-slate-800">
            <span class="flex h-6 w-6 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white">✓</span>
            <span class="text-slate-700 dark:text-slate-200">
              {{ ingredient.quantity ? `${ingredient.quantity} ${ingredient.unit || ''}` : '' }} {{ ingredient.name || ingredient.ingredient?.name }}
            </span>
          </li>
        </ul>
      </div>

      <div class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-2xl font-black text-slate-900 dark:text-white">Quick facts</h2>
        <div class="mt-5 space-y-4 text-sm text-slate-600 dark:text-slate-300">
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Difficulty</span><span class="font-semibold">{{ recipe.difficulty || 'Medium' }}</span></div>
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Servings</span><span class="font-semibold">{{ recipe.servings || 0 }}</span></div>
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Published</span><span class="font-semibold">{{ recipe.is_published ? 'Yes' : 'No' }}</span></div>
          <div class="flex items-center justify-between"><span>Price</span><span class="font-semibold">Free</span></div>
        </div>
      </div>
    </section>

    <!-- Comments Section -->
    <section v-if="recipe" class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <h2 class="text-2xl font-black text-slate-900 dark:text-white">Comments</h2>
      
      <div v-if="auth.isLoggedIn" class="mt-6">
        <textarea 
          v-model="newComment" 
          rows="3" 
          placeholder="Share your thoughts about this recipe..." 
          class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white"
        ></textarea>
        <button 
          @click="addComment" 
          :disabled="commentLoading || !newComment.trim()"
          class="mt-3 rounded-full bg-orange-500 px-6 py-2 text-sm font-semibold text-white transition hover:bg-orange-600 disabled:cursor-not-allowed disabled:opacity-60"
        >
          {{ commentLoading ? 'Posting...' : 'Post Comment' }}
        </button>
      </div>
      
      <div v-else class="mt-6 rounded-2xl bg-slate-50 p-4 text-center dark:bg-slate-800">
        <p class="text-sm text-slate-600 dark:text-slate-300">
          <NuxtLink to="/login" class="font-semibold text-orange-500">Login</NuxtLink> to leave a comment
        </p>
      </div>

      <div class="mt-8 space-y-4">
        <div v-if="comments.length === 0" class="text-center text-sm text-slate-500 dark:text-slate-400">
          No comments yet. Be the first to share your thoughts!
        </div>
        <div 
          v-for="comment in comments" 
          :key="comment.id" 
          class="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800"
        >
          <div class="flex items-center gap-3">
            <div class="flex h-10 w-10 items-center justify-center rounded-full bg-orange-500 text-sm font-bold text-white">
              {{ comment.user?.username?.charAt(0) || 'U' }}
            </div>
            <div>
              <div class="font-semibold text-slate-900 dark:text-white">{{ comment.user?.username || 'User' }}</div>
              <div class="text-xs text-slate-500 dark:text-slate-400">{{ new Date(comment.created_at).toLocaleDateString() }}</div>
            </div>
          </div>
          <p class="mt-3 text-sm text-slate-700 dark:text-slate-200">{{ comment.content }}</p>
        </div>
      </div>
    </section>
  </div>
</template>
