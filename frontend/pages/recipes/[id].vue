<script setup lang="ts">
const recipe = {
  title: 'Doro Wot',
  category: 'Ethiopian',
  creator: 'Fitsum',
  image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=1600&q=80',
  prepTime: 45,
  cookTime: 30,
  servings: 4,
  rating: 4.8,
  likes: 245,
  description: 'A rich Ethiopian chicken stew with berbere, onions and aromatic spices. Comforting, warm, and deeply flavorful.',
  ingredients: ['2 kg chicken', '3 tbsp berbere', '2 onions', '3 cloves garlic', '1 cup water', '2 tbsp butter'],
  steps: [
    'Heat the butter and sauté the onions until golden and fragrant.',
    'Add the garlic and berbere, then cook until the spices bloom.',
    'Add the chicken and stir well to coat every piece in the sauce.',
    'Simmer gently with water until tender and richly seasoned.',
    'Serve with injera and enjoy warm.'
  ]
}
</script>

<template>
  <div class="space-y-8">
    <section class="overflow-hidden rounded-[32px] border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <img :src="recipe.image" :alt="recipe.title" class="h-[420px] w-full object-cover" />
      <div class="p-6 md:p-8">
        <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <div class="mb-3 inline-flex rounded-full bg-orange-100 px-3 py-1 text-xs font-bold uppercase tracking-[0.18em] text-orange-600 dark:bg-orange-500/10 dark:text-orange-300">
              {{ recipe.category }}
            </div>
            <h1 class="text-4xl font-black text-slate-900 dark:text-white">{{ recipe.title }}</h1>
          </div>

          <div class="flex gap-3">
            <button class="rounded-full border border-slate-200 px-4 py-2 text-sm font-semibold text-slate-700 dark:border-slate-700 dark:text-slate-200">❤️ Save</button>
            <button class="rounded-full bg-orange-500 px-4 py-2 text-sm font-semibold text-white">👨‍🍳 Start Cooking</button>
          </div>
        </div>

        <div class="mt-6 flex flex-wrap items-center gap-4 text-sm text-slate-600 dark:text-slate-300">
          <span>👨‍🍳 {{ recipe.creator }}</span>
          <span>⏱ {{ recipe.prepTime }} min prep</span>
          <span>🍳 {{ recipe.cookTime }} min cook</span>
          <span>⭐ {{ recipe.rating }} / 5</span>
          <span>❤️ {{ recipe.likes }} likes</span>
        </div>

        <p class="mt-6 max-w-3xl text-lg text-slate-600 dark:text-slate-300">
          {{ recipe.description }}
        </p>
      </div>
    </section>

    <section class="grid gap-8 lg:grid-cols-[1fr_0.8fr]">
      <div class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-2xl font-black text-slate-900 dark:text-white">Ingredients</h2>
        <ul class="mt-5 space-y-3">
          <li v-for="ingredient in recipe.ingredients" :key="ingredient" class="flex items-center gap-3 rounded-2xl bg-slate-50 px-4 py-3 dark:bg-slate-800">
            <span class="flex h-6 w-6 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white">✓</span>
            <span class="text-slate-700 dark:text-slate-200">{{ ingredient }}</span>
          </li>
        </ul>
      </div>

      <div class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <h2 class="text-2xl font-black text-slate-900 dark:text-white">Quick facts</h2>
        <div class="mt-5 space-y-4 text-sm text-slate-600 dark:text-slate-300">
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Difficulty</span><span class="font-semibold">Medium</span></div>
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Servings</span><span class="font-semibold">{{ recipe.servings }}</span></div>
          <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800"><span>Cook mode</span><span class="font-semibold">Enabled</span></div>
          <div class="flex items-center justify-between"><span>Price</span><span class="font-semibold">Free</span></div>
        </div>
      </div>
    </section>

    <section class="rounded-[28px] border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <h2 class="text-2xl font-black text-slate-900 dark:text-white">Steps</h2>
      <ol class="mt-5 space-y-4">
        <li v-for="(step, index) in recipe.steps" :key="step" class="flex gap-4 rounded-2xl bg-slate-50 p-4 dark:bg-slate-800">
          <div class="flex h-8 w-8 items-center justify-center rounded-full bg-orange-500 text-sm font-bold text-white">{{ index + 1 }}</div>
          <p class="text-slate-700 dark:text-slate-200">{{ step }}</p>
        </li>
      </ol>
    </section>
  </div>
</template>
