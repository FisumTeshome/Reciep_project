// https://nuxt.com/docs/api/configuration/nuxt-config
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ['~/styles/globals.css'],
  modules: ['@nuxtjs/tailwindcss'],
  tailwindcss: {
    viewer: false,
  },
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true }
})
