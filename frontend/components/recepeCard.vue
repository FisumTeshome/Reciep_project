<script setup lang="ts">
interface RecipeCardProps {
  id?: string
  title: string
  image: string
  creator: string
  prepTime: string
  rating: number
  likes: number
  category: string
  description?: string
}

const { id, title, image, creator, prepTime, rating, likes, category, description = 'Freshly cooked and made to impress.' } = defineProps<RecipeCardProps>();
</script>

<template>
  <article class="group overflow-hidden rounded-[28px] border border-slate-200 bg-white shadow-sm transition duration-300 hover:-translate-y-1 hover:shadow-xl dark:border-slate-800 dark:bg-slate-900">
    <div class="relative overflow-hidden">
      <img :src="image" :alt="title" class="h-60 w-full object-cover transition duration-500 group-hover:scale-105" />
        <div class="absolute left-4 top-4 rounded-full bg-white/90 px-3 py-1 text-xs font-semibold text-slate-700 backdrop-blur dark:bg-slate-900/80 dark:text-slate-200">
        {{ category }}
      </div>
    </div>

    <div class="space-y-4 p-5">
      <div class="flex items-start justify-between gap-4">
        <div>
          <h3 class="text-xl font-bold text-slate-900 dark:text-white">{{ title }}</h3>
          <p class="mt-1 text-sm text-slate-500 dark:text-slate-400">{{ description }}</p>
        </div>
        <div class="rounded-full bg-amber-50 px-2 py-1 text-xs font-bold text-amber-700 dark:bg-amber-500/10 dark:text-amber-300">
          ★ {{ rating.toFixed(1) }}
        </div>
      </div>

      <div class="flex items-center justify-between text-sm text-slate-500 dark:text-slate-400">
        <span class="flex items-center gap-2">
          <span>⏱</span>
          {{ prepTime }}
        </span>
        <span class="flex items-center gap-2">
          <span>👨‍🍳</span>
          {{ creator }}
        </span>
      </div>

      <div class="flex items-center justify-between border-t border-slate-200 pt-4 dark:border-slate-800">
        <div class="flex items-center gap-2 text-sm text-slate-500 dark:text-slate-400">
          <span>❤️</span>
          {{ likes }}
        </div>
        <NuxtLink :to="`/recipes/${id || '1'}`" class="inline-flex items-center rounded-full bg-slate-900 px-4 py-2 text-sm font-semibold text-white transition hover:bg-orange-500 dark:bg-white dark:text-slate-900 dark:hover:bg-orange-500 dark:hover:text-white">
          View Recipe
        </NuxtLink>
      </div>
    </div>
  </article>
</template>
