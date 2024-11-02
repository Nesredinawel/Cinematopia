import { defineNuxtPlugin } from '#app';
import { ApolloProvider } from '@apollo/client';
import client from './client.js';

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(ApolloProvider, {
    defaultClient: client,
  });
});
