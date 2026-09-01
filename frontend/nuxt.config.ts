// https://nuxt.com/docs/api/configuration/nuxt-config
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ['~/styles/globals.css'],
  modules: ['@nuxtjs/tailwindcss'],
  tailwindcss: {
    viewer: false,
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_URL || 'http://localhost:8080',
      hasuraGraphql: process.env.NUXT_PUBLIC_HASURA_URL || 'http://localhost:8081',
    },
  },
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  vite: {
    server: {
      hmr: {
        protocol: 'ws',
        host: 'localhost',
        port: 3000,
      },
    },
  },
  ssr: true,
  nitro: {
    prerender: {
      crawlLinks: false,
    },
  },
  experimental: {
    payloadExtraction: false,
  },
})
