<template>
  <footer class="border-t border-slate-200 bg-white/80 dark:border-slate-800 dark:bg-slate-950/80">
    <div class="mx-auto flex max-w-7xl flex-col gap-8 px-4 py-10 text-sm text-slate-600 sm:px-6 lg:flex-row lg:items-center lg:justify-between lg:px-8 dark:text-slate-300">
      <div>
        <div class="text-lg font-black text-slate-900 dark:text-white">RecipeJoy</div>
        <p class="mt-2 max-w-md text-slate-500 dark:text-slate-400">
          Discover, save, cook and share unforgettable meals from around the world.
        </p>
      </div>

      <div class="flex flex-wrap items-center gap-6">
        <NuxtLink to="/" class="transition hover:text-orange-500">Home</NuxtLink>
        <NuxtLink to="/recipes" class="transition hover:text-orange-500">Recipes</NuxtLink>
        <NuxtLink to="/profile" class="transition hover:text-orange-500">Profile</NuxtLink>
      </div>
    </div>
  </footer>
</template>
