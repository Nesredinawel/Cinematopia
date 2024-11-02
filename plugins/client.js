import { defineNuxtPlugin } from '#app';
import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: process.env.HASURA_GRAPHQL_ENDPOINT || 'http://localhost:8080/v1/graphql',
  });

  const authLink = setContext((_, { headers }) => {
    const token = localStorage.getItem('token'); // For JWT
    return {
      headers: {
        ...headers,
        'x-hasura-admin-secret': process.env.HASURA_ADMIN_SECRET || 'mysecretkey',
        authorization: token ? `Bearer ${token}` : '',
      },
    };
  });

  const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  nuxtApp.provide('apollo', client);
});
