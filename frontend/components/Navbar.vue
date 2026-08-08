<template>
  <header class="sticky top-0 z-40 border-b border-slate-200/80 bg-white/80 backdrop-blur-xl dark:border-slate-800 dark:bg-slate-950/80">
    <div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-4 sm:px-6 lg:px-8">
      <NuxtLink to="/" class="flex items-center gap-3">
        <div class="flex h-10 w-10 items-center justify-center rounded-2xl bg-gradient-to-br from-orange-500 to-amber-400 shadow-lg shadow-orange-500/20">
          <span class="text-lg font-bold text-white">R</span>
        </div>
        <div>
          <div class="text-lg font-black tracking-tight text-slate-900 dark:text-white">RecipeJoy</div>
          <div class="text-[10px] uppercase tracking-[0.24em] text-slate-500 dark:text-slate-400">Cook & Share</div>
        </div>
      </NuxtLink>

      <nav class="hidden items-center gap-8 text-sm font-medium text-slate-600 md:flex dark:text-slate-300">
        <NuxtLink to="/" class="transition hover:text-orange-500">Home</NuxtLink>
        <NuxtLink to="/recipes" class="transition hover:text-orange-500">Recipes</NuxtLink>
        <NuxtLink to="/categories" class="transition hover:text-orange-500">Categories</NuxtLink>
        <NuxtLink to="/profile" class="transition hover:text-orange-500">Profile</NuxtLink>
      </nav>

      <div class="flex items-center gap-3">
        <button class="hidden rounded-full border border-slate-200 bg-slate-100 px-4 py-2 text-sm font-medium text-slate-700 transition hover:border-orange-200 hover:text-orange-600 sm:inline-flex dark:border-slate-700 dark:bg-slate-900 dark:text-slate-200">
          Dark mode
        </button>
        <NuxtLink to="/recipes" class="inline-flex items-center rounded-full bg-gradient-to-r from-orange-500 to-amber-500 px-4 py-2 text-sm font-semibold text-white shadow-lg shadow-orange-500/25 transition hover:opacity-95">
          Explore recipes
        </NuxtLink>
      </div>
    </div>
  </header>
</template>
