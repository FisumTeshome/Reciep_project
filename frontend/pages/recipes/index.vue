<script setup lang="ts">
const recipes = [
  { title: 'Doro Wot', category: 'Ethiopian', creator: 'Fitsum', prepTime: '45 min', rating: 4.8, likes: 245, image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Lentil Stew', category: 'African', creator: 'Amina', prepTime: '35 min', rating: 4.9, likes: 310, image: 'https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Chechebsa', category: 'Breakfast', creator: 'Hana', prepTime: '25 min', rating: 4.7, likes: 176, image: 'https://images.unsplash.com/photo-1525351484163-7529414344d8?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Coconut Rice', category: 'Dinner', creator: 'Mekdes', prepTime: '50 min', rating: 4.6, likes: 198, image: 'https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Berry Cheesecake', category: 'Desserts', creator: 'Sara', prepTime: '90 min', rating: 4.9, likes: 420, image: 'https://images.unsplash.com/photo-1551024601-bec78aea704b?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Garden Salad', category: 'Vegetarian', creator: 'Nahom', prepTime: '20 min', rating: 4.5, likes: 164, image: 'https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?auto=format&fit=crop&w=1200&q=80' }
]
</script>

<template>
  <div class="space-y-8">
    <section class="rounded-[32px] bg-slate-900 px-6 py-8 text-white shadow-xl">
      <div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-300">Browse recipes</p>
          <h1 class="mt-2 text-4xl font-black">Find a recipe you love</h1>
        </div>
        <div class="rounded-full bg-white/10 px-4 py-2 text-sm text-slate-200">365 recipes available</div>
      </div>
    </section>

    <section class="grid gap-6 lg:grid-cols-[280px_1fr]">
      <aside class="rounded-[28px] border border-slate-200 bg-white p-5 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-xl font-black text-slate-900 dark:text-white">Filters</h2>
        <div class="mt-5 space-y-5 text-sm text-slate-600 dark:text-slate-300">
          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Category</label>
            <select class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800">
              <option>All</option>
              <option>Ethiopian</option>
              <option>Desserts</option>
              <option>Vegetarian</option>
            </select>
          </div>

          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Prep time</label>
            <select class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800">
              <option>Any</option>
              <option>Under 30 min</option>
              <option>Under 60 min</option>
              <option>60+ min</option>
            </select>
          </div>

          <div>
            <label class="mb-2 block font-semibold text-slate-700 dark:text-slate-200">Difficulty</label>
            <select class="w-full rounded-xl border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800">
              <option>All</option>
              <option>Easy</option>
              <option>Medium</option>
              <option>Hard</option>
            </select>
          </div>
        </div>
      </aside>

      <div class="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
        <RecepeCard
          v-for="recipe in recipes"
          :key="recipe.title"
          :title="recipe.title"
          :image="recipe.image"
          :creator="recipe.creator"
          :prep-time="recipe.prepTime"
          :rating="recipe.rating"
          :likes="recipe.likes"
          :category="recipe.category"
        />
      </div>
    </section>
  </div>
</template>
