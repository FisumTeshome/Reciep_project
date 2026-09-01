<script setup lang="ts">
const categories = [
  { name: 'Ethiopian', image: 'https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=800&q=80', color: 'bg-orange-500/10 text-orange-700' },
  { name: 'African', image: 'https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?auto=format&fit=crop&w=800&q=80', color: 'bg-amber-500/10 text-amber-700' },
  { name: 'Desserts', image: 'https://images.unsplash.com/photo-1551024601-bec78aea704b?auto=format&fit=crop&w=800&q=80', color: 'bg-pink-500/10 text-pink-700' },
  { name: 'Vegetarian', image: 'https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?auto=format&fit=crop&w=800&q=80', color: 'bg-emerald-500/10 text-emerald-700' },
  { name: 'Breakfast', image: 'https://images.unsplash.com/photo-1525351484163-7529414344d8?auto=format&fit=crop&w=800&q=80', color: 'bg-yellow-500/10 text-yellow-700' },
  { name: 'Dinner', image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80', color: 'bg-indigo-500/10 text-indigo-700' }
]

const recipes = [
  { title: 'Doro Wot', category: 'Ethiopian', creator: 'Fitsum', prepTime: '45 min', rating: 4.8, likes: 245, image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Lentil Stew', category: 'African', creator: 'Amina', prepTime: '35 min', rating: 4.9, likes: 310, image: 'https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Chechebsa', category: 'Breakfast', creator: 'Hana', prepTime: '25 min', rating: 4.7, likes: 176, image: 'https://images.unsplash.com/photo-1525351484163-7529414344d8?auto=format&fit=crop&w=1200&q=80' },
  { title: 'Coconut Rice', category: 'Dinner', creator: 'Mekdes', prepTime: '50 min', rating: 4.6, likes: 198, image: 'https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?auto=format&fit=crop&w=1200&q=80' }
]

const creators = [
  { name: 'Fitsum', recipes: 42, likes: '1.2K' },
  { name: 'Amina', recipes: 29, likes: '980' },
  { name: 'Hana', recipes: 18, likes: '740' },
  { name: 'Mekdes', recipes: 24, likes: '890' }
]
</script>

<template>
  <div class="space-y-12 pb-12">
    <section class="relative overflow-hidden rounded-[32px] bg-gradient-to-br from-orange-600 via-amber-500 to-rose-500 px-6 py-12 text-white shadow-2xl shadow-orange-500/25 sm:px-8 lg:px-12 lg:py-16">
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_top_right,_rgba(255,255,255,0.22),_transparent_35%)]" />
      <div class="relative grid items-center gap-10 lg:grid-cols-[1.2fr_0.8fr]">
        <div>
          <div class="mb-4 inline-flex items-center gap-2 rounded-full bg-white/15 px-3 py-1 text-xs font-semibold uppercase tracking-[0.2em] text-orange-50 backdrop-blur-sm">
            Fresh picks • trending now
          </div>
          <h1 class="max-w-xl text-4xl font-black tracking-tight sm:text-5xl lg:text-6xl">
            Discover Your Next Favorite Recipe
          </h1>
          <p class="mt-5 max-w-lg text-base text-orange-50/90 sm:text-lg">
            Explore delicious recipes, discover new flavors, and share your own creations with a vibrant food community.
          </p>

          <div class="mt-8 flex flex-col gap-4 sm:flex-row">
            <NuxtLink to="/recipes" class="inline-flex items-center justify-center rounded-full bg-white px-6 py-3 text-sm font-bold text-orange-600 transition hover:bg-orange-50">
              Explore Recipes
            </NuxtLink>
            <NuxtLink to="/profile" class="inline-flex items-center justify-center rounded-full border border-white/40 bg-white/10 px-6 py-3 text-sm font-bold text-white transition hover:bg-white/20">
              Share Your Recipe
            </NuxtLink>
          </div>

          <div class="mt-8 flex max-w-xl items-center gap-3 rounded-2xl bg-white/10 p-3 backdrop-blur-sm">
            <span class="text-xl">🔍</span>
            <input
              type="text"
              placeholder="Search recipes, ingredients or dishes..."
              class="w-full border-0 bg-transparent text-sm text-white placeholder:text-orange-100/80 focus:outline-none"
            />
          </div>
        </div>

        <div class="relative">
          <div class="overflow-hidden rounded-[28px] border border-white/20 bg-white/10 p-3 shadow-2xl backdrop-blur-sm">
            <img
              src="https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=1200&q=80"
              alt="Featured recipe"
              class="h-[420px] w-full rounded-[24px] object-cover"
            />
          </div>
        </div>
      </div>
    </section>

    <section>
      <div class="mb-6 flex items-end justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-500">Popular categories</p>
          <h2 class="mt-2 text-3xl font-black tracking-tight text-slate-900 dark:text-white">Cook by mood</h2>
        </div>
      </div>

      <div class="grid gap-5 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6">
        <div v-for="category in categories" :key="category.name" class="group overflow-hidden rounded-[24px] border border-slate-200 bg-white shadow-sm transition hover:-translate-y-1 hover:shadow-lg dark:border-slate-800 dark:bg-slate-900">
          <img :src="category.image" :alt="category.name" class="h-40 w-full object-cover transition duration-500 group-hover:scale-105" />
          <div class="flex items-center justify-between p-4">
            <span class="text-base font-bold text-slate-900 dark:text-white">{{ category.name }}</span>
            <span :class="category.color" class="rounded-full px-2 py-1 text-xs font-semibold">Fresh</span>
          </div>
        </div>
      </div>
    </section>

    <section>
      <div class="mb-6 flex items-end justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-500">Trending recipes</p>
          <h2 class="mt-2 text-3xl font-black tracking-tight text-slate-900 dark:text-white">Most loved this week</h2>
        </div>
      </div>

      <div class="grid gap-6 lg:grid-cols-2 xl:grid-cols-4">
        <RecepeCard
          v-for="recipe in recipes"
          :key="recipe.title"
          id="1"
          :title="recipe.title"
          :image="recipe.image"
          :creator="recipe.creator"
          :prep-time="recipe.prepTime"
          :rating="recipe.rating"
          :likes="recipe.likes"
          :category="recipe.category"
          description="Comforting, vibrant and easy to cook at home."
        />
      </div>
    </section>

    <section class="rounded-[32px] bg-gradient-to-r from-amber-100 via-orange-50 to-rose-100 p-6 dark:from-slate-900 dark:via-slate-900 dark:to-slate-950 sm:p-8">
      <div class="flex flex-col gap-6 lg:flex-row lg:items-center lg:justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-600">Ethiopian Cuisine 🇪🇹</p>
          <h2 class="mt-2 text-3xl font-black tracking-tight text-slate-900 dark:text-white">Classic flavors from home</h2>
        </div>
        <NuxtLink to="/recipes" class="inline-flex items-center rounded-full bg-slate-900 px-5 py-3 text-sm font-bold text-white transition hover:bg-orange-500 dark:bg-white dark:text-slate-900 dark:hover:bg-orange-500 dark:hover:text-white">
          Explore Ethiopian dishes
        </NuxtLink>
      </div>

      <div class="mt-8 grid gap-4 md:grid-cols-2 xl:grid-cols-5">
        <div v-for="dish in ['Doro Wot', 'Tibs', 'Shiro', 'Kitfo', 'Firfir']" :key="dish" class="rounded-2xl border border-amber-200 bg-white/80 p-4 text-center shadow-sm dark:border-slate-700 dark:bg-slate-900/80">
          <div class="mb-3 text-3xl">🍲</div>
          <div class="font-bold text-slate-800 dark:text-white">{{ dish }}</div>
        </div>
      </div>
    </section>

    <section>
      <div class="mb-6 flex items-end justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.18em] text-orange-500">Featured creators</p>
          <h2 class="mt-2 text-3xl font-black tracking-tight text-slate-900 dark:text-white">Home cooks worth following</h2>
        </div>
      </div>

      <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-4">
        <div v-for="creator in creators" :key="creator.name" class="rounded-[24px] border border-slate-200 bg-white p-5 shadow-sm dark:border-slate-800 dark:bg-slate-900">
          <div class="flex items-center gap-4">
            <div class="flex h-14 w-14 items-center justify-center rounded-full bg-gradient-to-br from-orange-400 to-amber-300 text-lg font-black text-white">
              {{ creator.name.charAt(0) }}
            </div>
            <div>
              <div class="font-bold text-slate-900 dark:text-white">{{ creator.name }}</div>
              <div class="text-sm text-slate-500 dark:text-slate-400">{{ creator.recipes }} recipes</div>
            </div>
          </div>
          <div class="mt-4 text-sm text-slate-600 dark:text-slate-300">{{ creator.likes }} likes • Food enthusiast</div>
        </div>
      </div>
    </section>
  </div>
</template>
