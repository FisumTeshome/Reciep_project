<script setup lang="ts">
import { onMounted, ref, computed } from 'vue';
import { useRecipes } from '~/composables/useRecipes';

const { recipes, loading, error, fetchRecipes } = useRecipes();
const search = ref('');
const category = ref('');
const prepTime = ref('');
const recipesList = computed(() => recipes.value ?? []);

const loadRecipes = async () => {
  const params: Record<string, string> = {};
  if (search.value) params.search = search.value;
  if (category.value && category.value !== 'All') params.category_id = category.value;
  await fetchRecipes(params);
};

onMounted(() => {
  loadRecipes();
});
</script>

<template>
  <div class="space-y-8 px-4 py-12 sm:px-6 lg:px-8">
    <section class="rounded-[32px] bg-slate-900 px-6 py-8 text-white shadow-xl">
      <div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-300">Browse recipes</p>
          <h1 class="mt-2 text-4xl font-black">Find a recipe you love</h1>
        </div>
        <div class="rounded-full bg-white/10 px-4 py-2 text-sm text-slate-200">{{ recipesList.length }} recipes available</div>
      </div>

      <div class="mt-6 flex flex-col gap-3 sm:flex-row">
        <input v-model="search" @input="loadRecipes" type="text" placeholder="Search recipes, ingredients or dishes..." class="w-full rounded-2xl border border-slate-200 bg-white px-4 py-3 text-sm text-slate-900 outline-none focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
        <button @click="loadRecipes" class="rounded-full bg-orange-500 px-6 py-3 text-sm font-semibold text-white transition hover:bg-orange-600">Search</button>
      </div>
    </section>

    <section class="grid gap-6 lg:grid-cols-[280px_1fr]">
      <aside class="rounded-[28px] border border-slate-200 bg-white p-5 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-xl font-black text-slate-900 dark:text-white">Filters</h2>
        <div class="mt-5 space-y-5 text-sm text-slate-600 dark:text-slate-300">
          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Category</label>
            <select v-model="category" @change="loadRecipes" class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800 dark:text-white">
              <option>All</option>
              <option>Ethiopian</option>
              <option>Desserts</option>
              <option>Vegetarian</option>
            </select>
          </div>

          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Prep time</label>
            <select v-model="prepTime" class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800 dark:text-white" disabled>
              <option>Any</option>
              <option>Under 30 min</option>
              <option>Under 60 min</option>
              <option>60+ min</option>
            </select>
            <p class="mt-2 text-xs text-slate-500 dark:text-slate-400">Upcoming: client-side time filter.</p>
          </div>

          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Difficulty</label>
            <select class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800 dark:text-white" disabled>
              <option>All</option>
              <option>Easy</option>
              <option>Medium</option>
              <option>Hard</option>
            </select>
            <p class="mt-2 text-xs text-slate-500 dark:text-slate-400">Coming soon.</p>
          </div>
        </div>
      </aside>

      <div class="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
        <template v-if="loading">
          <div class="col-span-full rounded-[28px] bg-white p-8 text-center shadow-sm dark:bg-slate-900">Loading recipes...</div>
        </template>
        <template v-else-if="error">
          <div class="col-span-full rounded-[28px] bg-red-50 p-8 text-center text-red-700 shadow-sm dark:bg-red-900/20">{{ error }}</div>
        </template>
        <template v-else>
          <RecepeCard
            v-for="recipe in recipesList"
            :key="recipe.id || recipe.title"
            :id="recipe.id"
            :title="recipe.title"
            :image="recipe.featured_image || recipe.image || 'https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=1200&q=80'"
            :creator="recipe.user?.username || recipe.creator || 'Guest'"
            :prep-time="recipe.prep_time || recipe.prepTime"
            :rating="recipe.rating || 0"
            :likes="recipe.likes || recipe.like_count || 0"
            :category="recipe.category?.name || recipe.category"
          />
        </template>
      </div>
    </section>
  </div>
</template>
