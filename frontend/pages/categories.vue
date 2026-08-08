<script setup lang="ts">
const categories = [
  { name: 'Ethiopian', recipes: 134, image: 'https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=900&q=80' },
  { name: 'African', recipes: 89, image: 'https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?auto=format&fit=crop&w=900&q=80' },
  { name: 'Desserts', recipes: 64, image: 'https://images.unsplash.com/photo-1551024601-bec78aea704b?auto=format&fit=crop&w=900&q=80' },
  { name: 'Vegetarian', recipes: 72, image: 'https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?auto=format&fit=crop&w=900&q=80' },
  { name: 'Breakfast', recipes: 58, image: 'https://images.unsplash.com/photo-1525351484163-7529414344d8?auto=format&fit=crop&w=900&q=80' },
  { name: 'Dinner', recipes: 91, image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=900&q=80' }
]
</script>

<template>
  <div class="space-y-8">
    <section class="rounded-[32px] bg-slate-900 px-6 py-8 text-white shadow-xl">
      <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-300">Browse by category</p>
      <h1 class="mt-2 text-4xl font-black">Explore a world of food</h1>
    </section>

    <div class="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <div v-for="category in categories" :key="category.name" class="overflow-hidden rounded-[28px] border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <img :src="category.image" :alt="category.name" class="h-56 w-full object-cover" />
        <div class="p-5">
          <div class="flex items-center justify-between">
            <h2 class="text-2xl font-black text-slate-900 dark:text-white">{{ category.name }}</h2>
            <span class="rounded-full bg-orange-100 px-2 py-1 text-xs font-bold text-orange-700 dark:bg-orange-500/10 dark:text-orange-300">{{ category.recipes }} recipes</span>
          </div>
          <p class="mt-3 text-sm text-slate-600 dark:text-slate-300">Fresh inspiration, classic comfort food, and vibrant dishes from around the world.</p>
          <NuxtLink :to="`/recipes`" class="mt-5 inline-flex rounded-full bg-slate-900 px-4 py-2 text-sm font-semibold text-white transition hover:bg-orange-500 dark:bg-white dark:text-slate-900 dark:hover:bg-orange-500 dark:hover:text-white">
            View category
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
