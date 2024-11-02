export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',

  future: {
    compatibilityVersion: 4,
  },

  experimental: {
    sharedPrerenderData: false,
    compileTemplate: true,
    resetAsyncDataToUndefined: true,
    templateUtils: true,
    relativeWatchPaths: true,
    defaults: {
      useAsyncData: {
        deep: true,
      },
    },
  },

  unhead: {
    renderSSRHeadOptions: {
      omitLineBreaks: false,
    },
  },



  devtools: { enabled: true },

  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/google-fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxtjs/apollo', // Ensure this is installed
  ],

  googleFonts: {
    families: {
      Montserrat: true,
    },
  },

 

  apollo: {
    clients: {
      default: {
        httpEndpoint: 'http://localhost:8081', // Your Go backend
      },
    },
    // defaultOptions: {
    //   httpEndpoint: 'http://localhost:8081',
    // },
  },
  plugins: [
    './plugins/chapa.js'
  ],
})
