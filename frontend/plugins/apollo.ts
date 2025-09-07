import { ApolloClient, InMemoryCache } from '@apollo/client/core';
import { provideApolloClient } from '@vue/apollo-composable';

const apolloClient = new ApolloClient({
  uri: 'http://localhost:8080/v1/graphql', // Replace with your Hasura GraphQL endpoint
  cache: new InMemoryCache(),
});

export default defineNuxtPlugin(() => {
  provideApolloClient(apolloClient);
});
