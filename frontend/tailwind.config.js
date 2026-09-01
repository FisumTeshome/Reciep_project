/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.{vue,js,ts}',
    './pages/**/*.{vue,js,ts}',
    './composable/**/*.{js,ts,vue}',
    './plugins/**/*.{js,ts}',
    './app.vue',
    './app/**/**.{vue,js,ts}'
  ],
  darkMode: 'class',
  theme: {
    extend: {},
  },
  plugins: [],
}

