<template>
  <div class="space-y-8">
    <section class="overflow-hidden rounded-[32px] border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <div class="bg-gradient-to-r from-orange-500 via-amber-400 to-rose-400 p-8">
        <div class="flex flex-col gap-6 md:flex-row md:items-center md:justify-between">
          <div class="flex items-center gap-5">
            <div class="flex h-20 w-20 items-center justify-center rounded-full bg-white text-2xl font-black text-orange-600">
              F
            </div>
            <div class="text-white">
              <h1 class="text-3xl font-black">Fitsum</h1>
              <p class="text-orange-50">Food enthusiast 👨‍🍳</p>
            </div>
          </div>

          <button class="rounded-full bg-white px-5 py-3 text-sm font-bold text-orange-600 transition hover:bg-orange-50">
            Edit Profile
          </button>
        </div>
      </div>

      <div class="grid gap-4 p-6 md:grid-cols-3">
        <div class="rounded-2xl bg-slate-50 p-5 dark:bg-slate-800">
          <div class="text-sm text-slate-500 dark:text-slate-400">Recipes</div>
          <div class="mt-2 text-3xl font-black text-slate-900 dark:text-white">42</div>
        </div>
        <div class="rounded-2xl bg-slate-50 p-5 dark:bg-slate-800">
          <div class="text-sm text-slate-500 dark:text-slate-400">Likes</div>
          <div class="mt-2 text-3xl font-black text-slate-900 dark:text-white">1.2K</div>
        </div>
        <div class="rounded-2xl bg-slate-50 p-5 dark:bg-slate-800">
          <div class="text-sm text-slate-500 dark:text-slate-400">Saved</div>
          <div class="mt-2 text-3xl font-black text-slate-900 dark:text-white">128</div>
        </div>
      </div>
    </section>

    <section>
      <div class="grid gap-6 lg:grid-cols-3">
        <div class="rounded-[28px] border border-slate-200 bg-white p-5 shadow-sm dark:border-slate-800 dark:bg-slate-900 lg:col-span-2">
          <div class="mb-4 flex items-center justify-between">
            <h2 class="text-2xl font-black text-slate-900 dark:text-white">My Recipes</h2>
            <button class="rounded-full bg-orange-500 px-4 py-2 text-sm font-semibold text-white">New recipe</button>
          </div>
          <div class="grid gap-4 md:grid-cols-2">
            <div class="overflow-hidden rounded-2xl border border-slate-200 dark:border-slate-800">
              <img src="https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=900&q=80" class="h-40 w-full object-cover" />
              <div class="p-4">
                <div class="font-bold text-slate-900 dark:text-white">Doro Wot</div>
                <div class="mt-1 text-sm text-slate-500 dark:text-slate-400">45 min • Ethiopian</div>
              </div>
            </div>
            <div class="overflow-hidden rounded-2xl border border-slate-200 dark:border-slate-800">
              <img src="https://images.unsplash.com/photo-1547592180-85f173990554?auto=format&fit=crop&w=900&q=80" class="h-40 w-full object-cover" />
              <div class="p-4">
                <div class="font-bold text-slate-900 dark:text-white">Lentil Stew</div>
                <div class="mt-1 text-sm text-slate-500 dark:text-slate-400">35 min • African</div>
              </div>
            </div>
          </div>
        </div>

        <div class="rounded-[28px] border border-slate-200 bg-white p-5 shadow-sm dark:border-slate-800 dark:bg-slate-900">
          <h2 class="text-2xl font-black text-slate-900 dark:text-white">Saved</h2>
          <ul class="mt-5 space-y-4 text-sm text-slate-600 dark:text-slate-300">
            <li class="rounded-2xl bg-slate-50 p-3 dark:bg-slate-800">Goulash Bowl</li>
            <li class="rounded-2xl bg-slate-50 p-3 dark:bg-slate-800">Avocado Toast</li>
            <li class="rounded-2xl bg-slate-50 p-3 dark:bg-slate-800">Berry Cheesecake</li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>
